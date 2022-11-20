package apiError

type InstanceHookError struct {
	*BaseError
}

func NewInstanceHookError(code int, message string, data interface{}) InstanceHookError {
	return InstanceHookError{
		BaseError: NewBaseError(
			code,
			"INSTANCE_HOOK_ERROR",
			message,
			data,
		),
	}
}
