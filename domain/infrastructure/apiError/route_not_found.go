package apiError

type RouteNotFoundError struct {
	*BaseError
}

func NewRouteNotFoundError(data interface{}) *RouteNotFoundError {
	return &RouteNotFoundError{
		BaseError: NewBaseError(
			404,
			"ROUTE_NOT_FOUND",
			"Route not found",
			data,
		),
	}
}
