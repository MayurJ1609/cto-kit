package httpclient

import (
	"net/http"
)

const (
	defaultMaxRetryCount = 3
)

var DefaultClient = New(&http.Client{})

func New(client *http.Client, options ...Option) *http.Client {

	if client == nil {
		panic("client is nil")
	}

	nextReoundTrip := client.Transport
	if nextReoundTrip == nil {
		nextReoundTrip = http.DefaultTransport
	}

	retryRoundtrip := &Transport{
		Next:             nextReoundTrip,
		MaxRetryCount:    defaultMaxRetryCount,
		ShouldRetry:      defaultRetryPolicy,
		CalculateBackoff: DefaultBackoffPolicy,
	}

	for _, option := range options {
		option(retryRoundtrip)
	}

	return client
}
