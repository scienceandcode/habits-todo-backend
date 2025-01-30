package errors

type SimpleError struct {
	Message string `json:"message"`
}

func (e *SimpleError) Error() string {
	return e.Message
}

func NewSimpleError(message string) *SimpleError {
	return &SimpleError{
		Message: message,
	}
}
