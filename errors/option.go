package errors

type Option func(*option)

type option struct {
	message  string
	state    string
	errors   []error
	identity string
	service  Service
	data     interface{}
}

func WithMessage(message string) func(*option) {
	return func(o *option) {
		o.message = message
	}
}

func WithData(data interface{}) func(*option) {
	return func(o *option) {
		o.data = data
	}
}

func WithState(state string) func(*option) {
	return func(o *option) {
		o.state = state
	}
}

func WithTraceError(errors ...error) func(*option) {
	return func(o *option) {
		o.errors = errors
	}
}

func WithIdentity(identity string) func(*option) {
	return func(o *option) {
		o.identity = identity
	}
}

func WithService(service Service) func(*option) {
	return func(o *option) {
		o.service = service
	}
}
