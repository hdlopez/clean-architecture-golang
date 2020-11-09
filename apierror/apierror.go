package apierror

// New APIError instance
func New(code int, message string) *APIError {
	return &APIError{code, message}
}

func FromError(err error) *APIError {
	return New(500, err.Error())
}

// APIError is a trivial implementation of API errors response
type APIError struct {
	code    int
	message string
}

// Error function to be compliance with error interface
func (e *APIError) Error() string {
	return e.message
}

// Code funciton returns the error status code
func (e *APIError) Code() int {
	return e.code
}
