package apiError

type MethodNotSupportedError struct {
	*BaseError
}

func NewMethodNotSupportedError(data interface{}) *MethodNotSupportedError {
	return &MethodNotSupportedError{
		BaseError: NewBaseError(
			405,
			"METHOD_NOT_SUPPORTED",
			"Method not supported",
			data,
		),
	}
}
