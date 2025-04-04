package logging

import (
	"github.com/cto-kit/service"
)

type Option func(*option)

type option struct {
	reference   map[string]interface{}
	identity    string
	service     service.Service
	errors      []error
	appVersion  string
	prettyPrint bool
	stackTrace  bool
	referenceID string
}

func WithReference(reference map[string]interface{}) func(*option) {
	return func(o *option) {
		o.reference = reference
	}
}

func WithIdentity(identity string) func(*option) {
	return func(o *option) {
		o.identity = identity
	}
}

func WithService(service service.Service) func(*option) {
	return func(o *option) {
		o.service = service
	}
}

func WithTraceError(errors []error) func(*option) {
	return func(o *option) {
		o.errors = errors
	}
}

func WithAppVersion(appVersion string) func(*option) {
	return func(o *option) {
		o.appVersion = appVersion
	}
}

func WithPrettyPrint(prettyPrint bool) func(*option) {
	return func(o *option) {
		o.prettyPrint = prettyPrint
	}
}

func WithStackTrace(stackTrace bool) func(*option) {
	return func(o *option) {
		o.stackTrace = stackTrace
	}
}

func WithReferenceID(referenceID string) func(*option) {
	return func(o *option) {
		o.referenceID = referenceID
	}
}
