package vera

import (
	"regexp"
	"strings"
)

var errRegexp = map[error]*regexp.Regexp{
	ErrIncorrectCommandLineSpecified: regexp.MustCompile(`(?m)((?P<err>.*?)\.?\n)?Error: (?P<msg>Incorrect command line specified)\.`),
	ErrUnknownOption:                 regexp.MustCompile(`(?m)Error: (?P<msg>Unknown option: .*)`),
	ErrParameterIncorrect:            regexp.MustCompile(`(?m)Error: (?P<msg>Parameter incorrect: .*)`),
	ErrNoSuchVolumeMounted:           regexp.MustCompile(`(?m)Error: (?P<msg>No such volume is mounted)\.`),
	ErrAdministratorPrivileges:       regexp.MustCompile(`(?m)Error: (?P<msg>Failed to obtain administrator privileges)\.`),
	ErrNoSuchFileOrDirectory:         regexp.MustCompile(`(?m)Error: (?P<msg>No such file or directory):\n(?P<err>.*?)\n`),
	ErrNoVolumesMounted:              regexp.MustCompile(`(?m)Error: (?P<msg>No volumes mounted)\.`),
	ErrCommandNotFound:               regexp.MustCompile(`(?m)(?P<msg>.*: command not found)`),
	ErrOperationFailed:               regexp.MustCompile(`(?m)Error: (?P<msg>Operation failed) due to one or more of the following:`),
	ErrMountPointDoesNotExist:        regexp.MustCompile(`(?m)Error: mount: (?P<msg>.*: mount point does not exist)`),
	ErrMountPointIsNotADirectory:     regexp.MustCompile(`(?m)Error: mount: (?P<msg>.*: mount point is not a directory)`),
	ErrMountPointIsAlreadyInUse:      regexp.MustCompile(`(?m)Error: (?P<msg>Mount point is already in use)`),
	ErrVolumeAlreadyMounted:          regexp.MustCompile(`(?m)Error: (?P<msg>The volume you are trying to mount is already mounted)`),
}

type Error struct {
	cause   error  // the underlying error, e.g ErrNoSuchVolumeMounted
	message string // may contain additional information such as the slot number used
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

// parseError tries to match the passed stderr based on the errors contained in errRegexp map
func parseError(stderr string) Error {
	// iterate all known errors and try to match the stderr content
	for err, expr := range errRegexp {
		if expr.MatchString(stderr) {
			errMap := generateCaptureMap(stderr, expr)

			// the error does not contain additional details
			if expr.NumSubexp() == 1 {
				return Error{err, strings.ToLower(errMap["msg"])}
			}

			// unlike errors like ErrUnknownOption, ErrNoSuchFileOrDirectory contains only the path in the "err" capture
			// group. For this reason the actual error message is prefixed for this error
			if err == ErrNoSuchFileOrDirectory {
				return Error{err, strings.ToLower(errMap["msg"]) + ": " + errMap["err"]}
			}

			// the error provides more details, such as the ErrIncorrectCommandLineSpecified, which indicates, for
			// example, that a long option does not exist
			return Error{err, strings.ToLower(errMap["err"])}
		}
	}

	// capturing all errors returned by VeraCrypt is not that easy. Once most errors are mapped this error should be
	// the absolute exception
	return Error{ErrUnknown, ErrUnknown.Error()}
}
