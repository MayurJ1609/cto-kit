package config

type Option func(*option)
type option struct {
	defaultValue string
}

func WithDefault(defaultValue string) func(*option) {
	return func(o *option) {
		o.defaultValue = defaultValue
	}
}
