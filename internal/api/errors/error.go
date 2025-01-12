package errors

type Error struct {
	Message string        `json:"message"`
	Fields  []*FieldError `json:"fields"`
}

func (e *Error) Error() string {
	return e.Message
}

func NewError(message string, fields []*FieldError) *Error {
	return &Error{
		Message: message,
		Fields:  fields,
	}
}
