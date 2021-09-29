package vera

type Error struct {
	cause   error   // the underlying error, e.g ErrNoSuchVolumeMounted
	message string  // may contain additional information such as the slot number used
}

func (e Error) Error() string {
	return e.message
}

func (e Error) Unwrap() error {
	return e.cause
}

// newError returns an instance of Error with cause being the wrapped error and msg as the error message
func newError(cause error, msg string) *Error {
	return &Error{cause, msg}
}
