package internal

import (
	"context"
	"crypto/x509"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsRetryableStatusCode(t *testing.T) {
	retryable := []int{429, 500, 502, 503, 504}
	for _, code := range retryable {
		assert.True(t, isRetryableStatusCode(code), "expected %d to be retryable", code)
	}

	nonRetryable := []int{200, 201, 400, 401, 403, 404, 409, 422}
	for _, code := range nonRetryable {
		assert.False(t, isRetryableStatusCode(code), "expected %d to not be retryable", code)
	}
}

func TestParseRetryAfterSeconds(t *testing.T) {
	resp := &http.Response{Header: http.Header{"Retry-After": {"5"}}}
	assert.Equal(t, 5*time.Second, parseRetryAfter(resp))
}

func TestParseRetryAfterMissing(t *testing.T) {
	resp := &http.Response{Header: http.Header{}}
	assert.Equal(t, time.Duration(0), parseRetryAfter(resp))
}

func TestParseRetryAfterInvalid(t *testing.T) {
	resp := &http.Response{Header: http.Header{"Retry-After": {"invalid"}}}
	assert.Equal(t, time.Duration(0), parseRetryAfter(resp))
}

func TestRetryBackoff(t *testing.T) {
	minWait := 1 * time.Second
	maxWait := 30 * time.Second

	for attempt := 0; attempt < 5; attempt++ {
		wait := retryBackoff(attempt, minWait, maxWait)
		assert.GreaterOrEqual(t, wait, time.Duration(0), "backoff should not be negative")
		assert.LessOrEqual(t, wait, maxWait, "backoff should not exceed maxWait")
	}
}

func TestRetryBackoffExponentialGrowth(t *testing.T) {
	minWait := 1 * time.Second
	maxWait := 1 * time.Minute

	samples := 100
	var avg0, avg1, avg2 float64
	for i := 0; i < samples; i++ {
		avg0 += float64(retryBackoff(0, minWait, maxWait))
		avg1 += float64(retryBackoff(1, minWait, maxWait))
		avg2 += float64(retryBackoff(2, minWait, maxWait))
	}
	avg0 /= float64(samples)
	avg1 /= float64(samples)
	avg2 /= float64(samples)

	assert.Greater(t, avg1, avg0, "attempt 1 should have longer average wait than attempt 0")
	assert.Greater(t, avg2, avg1, "attempt 2 should have longer average wait than attempt 1")
}

func TestCallAPIRetryOn429(t *testing.T) {
	var callCount atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		n := callCount.Add(1)
		if n <= 2 {
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	cfg := &Configuration{
		HTTPClient: server.Client(),
		RetryConfig: &RetryConfig{
			MaxRetries: 3,
			MinWait:    10 * time.Millisecond,
			MaxWait:    50 * time.Millisecond,
		},
	}
	client := &APIClient{Cfg: cfg}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, server.URL, nil)
	require.NoError(t, err)

	resp, err := client.CallAPI(req)
	require.NoError(t, err)
	resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, int32(3), callCount.Load())
}

func TestCallAPIRetryOn500(t *testing.T) {
	var callCount atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		n := callCount.Add(1)
		if n == 1 {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	cfg := &Configuration{
		HTTPClient: server.Client(),
		RetryConfig: &RetryConfig{
			MaxRetries: 3,
			MinWait:    10 * time.Millisecond,
			MaxWait:    50 * time.Millisecond,
		},
	}
	client := &APIClient{Cfg: cfg}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, server.URL, nil)
	require.NoError(t, err)

	resp, err := client.CallAPI(req)
	require.NoError(t, err)
	resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, int32(2), callCount.Load())
}

func TestCallAPINoRetryOn400(t *testing.T) {
	var callCount atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callCount.Add(1)
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer server.Close()

	cfg := &Configuration{
		HTTPClient: server.Client(),
		RetryConfig: &RetryConfig{
			MaxRetries: 3,
			MinWait:    10 * time.Millisecond,
			MaxWait:    50 * time.Millisecond,
		},
	}
	client := &APIClient{Cfg: cfg}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, server.URL, nil)
	require.NoError(t, err)

	resp, err := client.CallAPI(req)
	require.NoError(t, err)
	resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	assert.Equal(t, int32(1), callCount.Load(), "should not retry on 400")
}

func TestCallAPIRetryExhausted(t *testing.T) {
	var callCount atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callCount.Add(1)
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	cfg := &Configuration{
		HTTPClient: server.Client(),
		RetryConfig: &RetryConfig{
			MaxRetries: 2,
			MinWait:    10 * time.Millisecond,
			MaxWait:    50 * time.Millisecond,
		},
	}
	client := &APIClient{Cfg: cfg}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, server.URL, nil)
	require.NoError(t, err)

	resp, err := client.CallAPI(req)
	require.NoError(t, err)
	resp.Body.Close()

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	assert.Equal(t, int32(3), callCount.Load(), "should be 1 initial + 2 retries")
}

func TestCallAPIRetryPreservesBody(t *testing.T) {
	var bodies []string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, _ := readBody(r)
		bodies = append(bodies, string(b))
		if len(bodies) < 2 {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	cfg := &Configuration{
		HTTPClient: server.Client(),
		RetryConfig: &RetryConfig{
			MaxRetries: 3,
			MinWait:    10 * time.Millisecond,
			MaxWait:    50 * time.Millisecond,
		},
	}
	client := &APIClient{Cfg: cfg}

	body := `{"name":"test"}`
	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, server.URL, strings.NewReader(body))
	require.NoError(t, err)

	resp, err := client.CallAPI(req)
	require.NoError(t, err)
	resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	require.Len(t, bodies, 2)
	assert.Equal(t, body, bodies[0], "first attempt body should match")
	assert.Equal(t, body, bodies[1], "retry body should match original")
}

func readBody(r *http.Request) ([]byte, error) {
	if r.Body == nil {
		return nil, nil
	}
	defer r.Body.Close()
	b := make([]byte, r.ContentLength)
	_, err := r.Body.Read(b)
	return b, err
}

func TestCallAPIRetryContextCancellation(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	cfg := &Configuration{
		HTTPClient: server.Client(),
		RetryConfig: &RetryConfig{
			MaxRetries: 10,
			MinWait:    1 * time.Second,
			MaxWait:    5 * time.Second,
		},
	}
	client := &APIClient{Cfg: cfg}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, server.URL, nil)
	require.NoError(t, err)

	_, err = client.CallAPI(req)
	assert.Error(t, err, "should fail due to context cancellation during backoff")
}

func TestCallAPIRetryWithRetryAfterHeader(t *testing.T) {
	var callCount atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		n := callCount.Add(1)
		if n == 1 {
			w.Header().Set("Retry-After", "1")
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	cfg := &Configuration{
		HTTPClient: server.Client(),
		RetryConfig: &RetryConfig{
			MaxRetries: 3,
			MinWait:    10 * time.Millisecond,
			MaxWait:    5 * time.Second,
		},
	}
	client := &APIClient{Cfg: cfg}

	start := time.Now()
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, server.URL, nil)
	require.NoError(t, err)

	resp, err := client.CallAPI(req)
	require.NoError(t, err)
	resp.Body.Close()

	elapsed := time.Since(start)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.GreaterOrEqual(t, elapsed, 900*time.Millisecond, "should have respected Retry-After: 1")
}

func TestCallAPINoRetryWhenDisabled(t *testing.T) {
	var callCount atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callCount.Add(1)
		w.WriteHeader(http.StatusInternalServerError)
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

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	assert.Equal(t, int32(1), callCount.Load(), "should not retry when RetryConfig is nil")
}

func TestNewConfigurationDefaultRetry(t *testing.T) {
	cfg := NewConfiguration()
	require.NotNil(t, cfg.RetryConfig, "retry should be enabled by default")
	assert.Equal(t, DefaultMaxRetries, cfg.RetryConfig.MaxRetries)
	assert.Equal(t, DefaultRetryMinWait, cfg.RetryConfig.MinWait)
	assert.Equal(t, DefaultRetryMaxWait, cfg.RetryConfig.MaxWait)
}

func TestNewConfigurationRetryFromEnv(t *testing.T) {
	t.Setenv("INFOBLOX_MAX_RETRIES", "5")
	t.Setenv("INFOBLOX_RETRY_MIN_WAIT", "2s")
	t.Setenv("INFOBLOX_RETRY_MAX_WAIT", "1m")
	cfg := NewConfiguration()
	require.NotNil(t, cfg.RetryConfig)
	assert.Equal(t, 5, cfg.RetryConfig.MaxRetries)
	assert.Equal(t, 2*time.Second, cfg.RetryConfig.MinWait)
	assert.Equal(t, 1*time.Minute, cfg.RetryConfig.MaxWait)
}

func TestNewConfigurationRetryDisabledByEnv(t *testing.T) {
	t.Setenv("INFOBLOX_MAX_RETRIES", "0")
	cfg := NewConfiguration()
	assert.Nil(t, cfg.RetryConfig, "setting retry attempts to 0 should disable retry")
}

func TestIsRetryableError(t *testing.T) {
	assert.False(t, isRetryableError(nil), "nil error should not be retryable")

	assert.True(t, isRetryableError(errors.New("connection reset")), "generic error should be retryable")

	dnsNotFound := &net.DNSError{Err: "no such host", Name: "bad.example.com", IsNotFound: true}
	assert.False(t, isRetryableError(dnsNotFound), "DNS NXDOMAIN should not be retryable")

	dnsTemp := &net.DNSError{Err: "temporary failure", Name: "example.com", IsTemporary: true}
	assert.True(t, isRetryableError(dnsTemp), "temporary DNS error should be retryable")

	certErr := &x509.CertificateInvalidError{Reason: x509.Expired}
	assert.False(t, isRetryableError(certErr), "certificate invalid error should not be retryable")

	unknownAuth := &x509.UnknownAuthorityError{}
	assert.False(t, isRetryableError(unknownAuth), "unknown authority error should not be retryable")

	wrappedCert := fmt.Errorf("tls: %w", &x509.CertificateInvalidError{Reason: x509.Expired})
	assert.False(t, isRetryableError(wrappedCert), "wrapped certificate error should not be retryable")
}

func TestCallAPINoRetryOnPermanentError(t *testing.T) {
	var callCount atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callCount.Add(1)
		w.WriteHeader(http.StatusOK)
	}))
	server.Close()

	cfg := &Configuration{
		HTTPClient: server.Client(),
		RetryConfig: &RetryConfig{
			MaxRetries: 3,
			MinWait:    10 * time.Millisecond,
			MaxWait:    50 * time.Millisecond,
		},
	}
	client := &APIClient{Cfg: cfg}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, server.URL, nil)
	require.NoError(t, err)

	_, err = client.CallAPI(req)
	assert.Error(t, err, "should fail on closed server")
	assert.Equal(t, int32(0), callCount.Load(), "should not reach the handler on a closed server")
}
