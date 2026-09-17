package option

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/infobloxopen/universal-ddi-go-client/internal"
)

func TestWithCSPUrl(t *testing.T) {
	config := &internal.Configuration{}
	url := "http://test.com"
	opt := WithCSPUrl(url)
	opt(config)
	assert.Equal(t, url, config.CSPURL)
}

func TestWithAPIKey(t *testing.T) {
	config := &internal.Configuration{}
	apiKey := "testKey"
	opt := WithAPIKey(apiKey)
	opt(config)
	assert.Equal(t, apiKey, config.APIKey)
}

func TestWithHTTPClient(t *testing.T) {
	config := &internal.Configuration{}
	client := &http.Client{}
	opt := WithHTTPClient(client)
	opt(config)
	assert.Equal(t, client, config.HTTPClient)
}

func TestWithDefaultTags(t *testing.T) {
	config := &internal.Configuration{}
	tags := map[string]string{"tag1": "value1"}
	opt := WithDefaultTags(tags)
	opt(config)
	assert.Equal(t, tags, config.DefaultTags)
}

func TestWithClientName(t *testing.T) {
	config := &internal.Configuration{}
	name := "testClient"
	opt := WithClientName(name)
	opt(config)
	assert.Equal(t, name, config.ClientName)
}

func TestWithDebug(t *testing.T) {
	config := &internal.Configuration{}
	opt := WithDebug(true)
	opt(config)
	assert.Equal(t, true, config.Debug)
}

func TestWithRateLimit(t *testing.T) {
	config1 := &internal.Configuration{}
	config2 := &internal.Configuration{}
	opt := WithRateLimit(10, 5)
	opt(config1)
	opt(config2)
	assert.NotNil(t, config1.RateLimiter)
	assert.NotNil(t, config2.RateLimiter)
	assert.Same(t, config1.RateLimiter, config2.RateLimiter, "same limiter instance should be shared")
}

func TestWithRateLimitDisabled(t *testing.T) {
	config := &internal.Configuration{}
	config.RateLimiter = &mockRateLimiter{}
	opt := WithRateLimitDisabled()
	opt(config)
	assert.Nil(t, config.RateLimiter)
}

type mockRateLimiter struct{}

func (m *mockRateLimiter) Wait(ctx context.Context) error { return nil }

func TestWithRateLimiter(t *testing.T) {
	config := &internal.Configuration{}
	limiter := &mockRateLimiter{}
	opt := WithRateLimiter(limiter)
	opt(config)
	assert.Equal(t, limiter, config.RateLimiter)
}

func TestWithRateLimiterNil(t *testing.T) {
	config := &internal.Configuration{
		RateLimiter: &mockRateLimiter{},
	}
	opt := WithRateLimiter(nil)
	opt(config)
	assert.NotNil(t, config.RateLimiter, "WithRateLimiter(nil) should not clear an existing limiter")
}

func TestWithRateLimitZeroDisables(t *testing.T) {
	config := &internal.Configuration{}
	opt := WithRateLimit(0, 10)
	opt(config)
	assert.Nil(t, config.RateLimiter, "rate <= 0 should disable the rate limiter")
}

func TestWithRateLimitNegativeDisables(t *testing.T) {
	config := &internal.Configuration{}
	opt := WithRateLimit(-5, 10)
	opt(config)
	assert.Nil(t, config.RateLimiter, "negative rate should disable the rate limiter")
}

func TestWithRateLimitBurstClampsToOne(t *testing.T) {
	config := &internal.Configuration{}
	opt := WithRateLimit(10, 0)
	opt(config)
	assert.NotNil(t, config.RateLimiter, "burst 0 should be clamped to 1, not cause failure")
}

func TestWithRetry(t *testing.T) {
	config := &internal.Configuration{}
	opt := WithRetry(5, 2*time.Second, 1*time.Minute)
	opt(config)
	require.NotNil(t, config.RetryConfig)
	assert.Equal(t, 5, config.RetryConfig.MaxRetries)
	assert.Equal(t, 2*time.Second, config.RetryConfig.MinWait)
	assert.Equal(t, 1*time.Minute, config.RetryConfig.MaxWait)
}

func TestWithRetryZeroDisables(t *testing.T) {
	config := &internal.Configuration{
		RetryConfig: &internal.RetryConfig{MaxRetries: 3},
	}
	opt := WithRetry(0, 0, 0)
	opt(config)
	assert.Nil(t, config.RetryConfig, "zero attempts should disable retry")
}

func TestWithRetryDisabled(t *testing.T) {
	config := &internal.Configuration{
		RetryConfig: &internal.RetryConfig{MaxRetries: 3},
	}
	opt := WithRetryDisabled()
	opt(config)
	assert.Nil(t, config.RetryConfig)
}
