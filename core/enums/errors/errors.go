package errors

const (
	ValidationError           = 1
	AuthenticationError       = 2
	AuthorizationError        = 3
	InternalServerError       = 4
	ValidationErrorStatus     = "Validation Error"
	AuthenticationErrorStatus = "Authentication Error"
	AuthorizationErrorStatus  = "Authorization Error"
	InternalServerErrorStatus = "Internal Server Error"
)
