package token

import "time"

// Option is a function adapter to change config of the errors struct.
type Option func(*option)

type option struct {
	expire time.Duration
}

// WithDuration injects the duration
func WithDuration(expire time.Duration) func(*option) {
	return func(o *option) {
		if expire != 0 {
			o.expire = expire
		}
	}
}
