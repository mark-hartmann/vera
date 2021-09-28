package vera

import (
	"errors"
)

var ErrNoVolumePath = errors.New("no volume path provided")
var ErrNoVolumesMounted = errors.New("no volumes mounted")
var ErrNoSuchVolumeMounted = errors.New("no such value is mounted")
var ErrParameterIncorrect = errors.New("parameter incorrect")

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

func newError(cause error, msg string) *Error {
	return &Error{cause, msg}
}
