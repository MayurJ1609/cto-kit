package apm

type Option func(*option)

type option struct {
	agentType   agentType
	enable      bool
	serviceName string
	secretToken string
	serverURL   string
}

func WithServerURL(url string) Option {
	return func(o *option) {
		o.secretToken = url
	}
}

func WithServiceName(name string) Option {
	return func(o *option) {
		o.serviceName = name
	}
}

func WithServiceToken(token string) Option {
	return func(o *option) {
		o.secretToken = token
	}
}

func WithAgentType(agentType agentType) Option {
	return func(o *option) {
		o.agentType = agentType
	}
}

func WithEnableMonitoring(enable bool) Option {
	return func(o *option) {
		o.enable = enable
	}
}
