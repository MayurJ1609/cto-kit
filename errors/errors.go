package errors

type (
	// Error is a custom error type that can be used to wrap errors

	Error interface {
		// Errorcode and must have values defines by package
		Code() string

		// Error message
		Error() string

		// Underlying error
		Data() interface{}
	}

	appError struct {
		ErrorCode        string      `json:"error,omitempty"`
		Message          string      `json:"message,omitempty"`
		ErrorDescription string      `json:"errorDescription,omitempty"`
		ErrorState       string      `json:"state,omitempty"`
		Service          Service     `json:"service,omitempty"`
		Identity         string      `json:"identity,omitempty"`
		AdditionalData   interface{} `json:"additionalData,omitempty"`
		Errors           []error     `json:"errors,omitempty"`
	}
)

// New returns a new Error
func New(code string, description string, opts ...Option) Error {
	option := &option{}
	for _, opt := range opts {
		opt(option)
	}

	return &appError{
		ErrorCode:        code,
		Message:          option.message,
		ErrorDescription: description,
		ErrorState:       option.state,
		Service:          option.service,
		Identity:         option.identity,
		AdditionalData:   option.data,
		Errors:           option.errors,
	}

}

func (e *appError) Code() string {
	return e.ErrorCode
}

func (e *appError) Error() string {
	return e.ErrorDescription
}

func (e *appError) Data() interface{} {
	return e.AdditionalData
}
