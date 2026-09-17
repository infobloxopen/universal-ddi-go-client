package internal

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/time/rate"
)

func TestCallAPIWithRateLimiter(t *testing.T) {
	var callCount atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callCount.Add(1)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	cfg := &Configuration{
		HTTPClient:  server.Client(),
		RateLimiter: rate.NewLimiter(rate.Limit(100), 100),
	}
	client := &APIClient{Cfg: cfg}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, server.URL, nil)
	require.NoError(t, err)

	resp, err := client.CallAPI(req)
	require.NoError(t, err)
	resp.Body.Close()

	assert.Equal(t, int32(1), callCount.Load())
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestCallAPIWithoutRateLimiter(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	cfg := &Configuration{
		HTTPClient: server.Client(),
	}
	client := &APIClient{Cfg: cfg}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, server.URL, nil)
	require.NoError(t, err)

	resp, err := client.CallAPI(req)
	require.NoError(t, err)
	resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestCallAPIRateLimiterContextCancelled(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("HTTP call should not have been made")
	}))
	defer server.Close()

	cfg := &Configuration{
		HTTPClient:  server.Client(),
		RateLimiter: rate.NewLimiter(rate.Limit(0), 0),
	}
	client := &APIClient{Cfg: cfg}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, server.URL, nil)
	require.NoError(t, err)

	_, err = client.CallAPI(req)
	assert.Error(t, err)
}

func TestCallAPIRateLimiterThroughput(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	requestsPerSecond := 20.0
	cfg := &Configuration{
		HTTPClient:  server.Client(),
		RateLimiter: rate.NewLimiter(rate.Limit(requestsPerSecond), 1),
	}
	client := &APIClient{Cfg: cfg}

	totalRequests := 10
	start := time.Now()

	for i := 0; i < totalRequests; i++ {
		req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, server.URL, nil)
		require.NoError(t, err)

		resp, err := client.CallAPI(req)
		require.NoError(t, err)
		resp.Body.Close()
	}

	elapsed := time.Since(start)
	expectedMin := time.Duration(float64(totalRequests-1)/requestsPerSecond*1000) * time.Millisecond
	assert.GreaterOrEqual(t, elapsed, expectedMin-50*time.Millisecond,
		"requests completed too quickly, rate limiter may not be working")
}

func TestCallAPIRateLimiterSkipsGET(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	requestsPerSecond := 20.0
	cfg := &Configuration{
		HTTPClient:  server.Client(),
		RateLimiter: rate.NewLimiter(rate.Limit(requestsPerSecond), 1),
	}
	client := &APIClient{Cfg: cfg}

	totalRequests := 10
	start := time.Now()

	for i := 0; i < totalRequests; i++ {
		req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, server.URL, nil)
		require.NoError(t, err)

		resp, err := client.CallAPI(req)
		require.NoError(t, err)
		resp.Body.Close()
	}

	elapsed := time.Since(start)
	assert.Less(t, elapsed, 200*time.Millisecond,
		"GET requests should not be throttled by the rate limiter")
}

func TestNewConfigurationDefaultRateLimit(t *testing.T) {
	cfg := NewConfiguration()
	assert.NotNil(t, cfg.RateLimiter, "rate limiter should be enabled by default")
}

func TestNewConfigurationRateLimitFromEnv(t *testing.T) {
	t.Setenv("INFOBLOX_RATE_LIMIT", "10")
	cfg := NewConfiguration()
	assert.NotNil(t, cfg.RateLimiter)
}

func TestNewConfigurationRateLimitBurstFromEnv(t *testing.T) {
	t.Setenv("INFOBLOX_RATE_LIMIT", "10")
	t.Setenv("INFOBLOX_RATE_LIMIT_BURST", "5")
	cfg := NewConfiguration()
	assert.NotNil(t, cfg.RateLimiter)
}

func TestNewConfigurationRateLimitDisabledByEnv(t *testing.T) {
	t.Setenv("INFOBLOX_RATE_LIMIT", "0")
	cfg := NewConfiguration()
	assert.Nil(t, cfg.RateLimiter, "setting rate limit to 0 should disable it")
}

func TestNewConfigurationRateLimitInvalidEnvFallsToDefault(t *testing.T) {
	t.Setenv("INFOBLOX_RATE_LIMIT", "not-a-number")
	cfg := NewConfiguration()
	assert.NotNil(t, cfg.RateLimiter, "invalid env should fall back to default (enabled)")
}
