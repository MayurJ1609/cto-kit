package apm

import "errors"

// application monitoring error's
var (
	ErrInvalidAppName     = errors.New("application name is not set")
	ErrInvalidLicenseName = errors.New("APM License is not set")
	ErrUnsupported        = errors.New("unsupported monitoring system type")
	ErrNotInitialized     = errors.New("APM application object is nil")
	ErrInvalidTrans       = errors.New("APM transaction object is nil")
	ErrInvalidRequest     = errors.New("HTTP request object didn't provided")
	ErrInvalidSegment     = errors.New("APM segment object is nil")
	ErrInvalidDataType    = errors.New("invalid data type")
)
