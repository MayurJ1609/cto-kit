package httpclient

import (
	"crypto/x509"
	"net/http"
	"net/url"
	"strings"
)

type RetryPolicy func(statusCode int, err error) bool

var defaultRetryPolicy RetryPolicy = func(statusCode int, err error) bool {

	t, ok := err.(interface {
		Temporary() bool
	})
	if ok && t.Temporary() {
		return true
	}

	switch e := err.(type) {
	case *url.Error:
		switch {
		case
			e.Op == "parse",
			strings.Contains(e.Err.Error(), "stopped after"),
			strings.Contains(e.Err.Error(), "unsupported protocol scheme"),
			strings.Contains(e.Err.Error(), "no such host in request URL"):
			return false
		}

		switch e.Err.(type) {
		case
			x509.UnknownAuthorityError,
			x509.HostnameError,
			x509.CertificateInvalidError,
			x509.ConstraintViolationError:
			return false
		}
	case error:
		return true
	case nil:
	}

	switch statusCode {
	case
		http.StatusLocked,
		http.StatusInternalServerError,
		http.StatusBadGateway,
		http.StatusServiceUnavailable,
		http.StatusGatewayTimeout,
		http.StatusTooManyRequests,
		http.StatusInsufficientStorage:
		return true
	case 0:
		return true
	default:
		return false
	}
}
