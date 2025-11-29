package apierror

// New APIError instance
func New(code int, message string) *APIError {
	return &APIError{code, message}
}

// FromError creates a new APIError from a standard error
func FromError(err error) *APIError {
	return New(500, err.Error())
}

// APIError is a trivial implementation of API errors response
type APIError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// Error function to be compliance with error interface
func (e *APIError) Error() string {
	return e.Message
}

// Code funciton returns the error status code
func (e *APIError) Code() int {
	return e.Status
}
