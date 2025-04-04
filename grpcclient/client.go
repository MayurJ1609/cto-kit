package grpcclient

// this package is a wrapper around the grpc client to provide retry functionality
// for grpc calls. It is a wrapper around the grpc client and provides a new client
// that can be used to make grpc calls with retry functionality.
//
// The client is created using the New function. The New function takes a grpc client
// and a list of options. The options are used to configure the retry behavior of the
// client. The client returned by the New function can be used to make grpc calls with
// retry functionality.
//
// The client uses a backoff policy to determine the time to wait before retrying a
// failed call. The backoff policy is a function that takes the number of retries and
// returns the time to wait before retrying. The client uses the backoff policy to
// determine the time to wait before retrying a failed call.
//
// The client also uses a retry policy to determine whether to retry a failed call. The
// retry policy is a function that takes the error returned by the call and returns
// whether to retry the call. The client uses the retry policy to determine whether to
// retry a failed call.
//
// The client also uses a max retry count to determine the maximum number of times to
// retry a failed call. The client uses the max retry count to determine the maximum
// number of times to retry a failed call.
//
// The client also uses a next round trip to determine the next round trip to use for
// the call. The client uses the next round trip to determine the next round trip to
// use for the call.
//
// The client also uses a dial option to determine the dial options to use for the call.
// The client uses the dial option to determine the dial options to use for the call.
//
// The client also uses a call option to determine the call options to use for the call.
// The client uses the call option to determine the call options to use for the call.
//
// The client also uses a stream option to determine the stream options to use for the call.
// The client uses the stream option to determine the stream options to use for the call.
//
// The client also uses a unary interceptor to determine the unary interceptors to use for the call.

import (
	"google.golang.org/grpc"
)

const (
	defaultMaxRetryCount = 3
)

// Client is a wrapper around the grpc client to provide retry functionality
type Client struct {
	client           *grpc.ClientConn
	maxRetryCount    int
	shouldRetry      RetryPolicy
	calculateBackoff BackoffPolicy
	// nextRoundTrip     grpc.RoundRobin
	dialOption        []grpc.DialOption
	callOption        []grpc.CallOption
	streamOption      []grpc.CallOption
	unaryInterceptors []grpc.UnaryClientInterceptor
}

// New creates a new grpc client with retry functionality
func New(client *grpc.ClientConn, options ...Option) *Client {
	if client == nil {
		panic("client is nil")
	}

	retryClient := &Client{
		client:           client,
		maxRetryCount:    defaultMaxRetryCount,
		shouldRetry:      defaultRetryPolicy,
		calculateBackoff: DefaultBackoffPolicy,
		// nextRoundTrip:    grpc.RoundRobin{},
		dialOption:        []grpc.DialOption{},
		callOption:        []grpc.CallOption{},
		streamOption:      []grpc.CallOption{},
		unaryInterceptors: []grpc.UnaryClientInterceptor{},
	}

	for _, option := range options {
		option(retryClient)
	}

	return retryClient
}

// Option is a function that configures the retry client
type Option func(*Client)
