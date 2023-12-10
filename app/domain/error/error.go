package error

type Error struct {
	description string
}

func (e *Error) Error() string {
	return e.description
}

func NewError(description string) *Error {
	return &Error{
		description: description,
	}
}

var ErrNotFound = NewError("not found")
