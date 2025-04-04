package grpcclient

import (
	"net/http"
)

type Transport struct {
	// Transport is an http.RoundTripper that is used to make HTTP requests.
	// It is nil if the default transport should be used.
	Next             http.RoundTripper
	MaxRetryCount    int
	ShouldRetry      RetryPolicy
	CalculateBackoff BackoffPolicy
}
