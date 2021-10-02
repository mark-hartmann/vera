package vera

import (
	"errors"
)

// ErrIncorrectCommandLineSpecified gets returned when trying to use an option that does not exist in the currently
// installed VeraCrypt version
var ErrIncorrectCommandLineSpecified = errors.New("incorrect command line specified")

// ErrUnknownOption is returned when trying to pass an unknown value for a parameter which only accepts certain
// predefined values. An example would be --hash=unknown-algo
var ErrUnknownOption = errors.New("unknown option")

// ErrParameterIncorrect gets returned when using an invalid value for an option that requires a specific format such
// as digits only. An example would be --slot=one instead of --slot=1
var ErrParameterIncorrect = errors.New("parameter incorrect")

// ErrNoSuchVolumeMounted gets returned when trying to work on a volume that's not mounted. An example would be trying
// to dismount a volume at a slot that has no volume mounted, e.g. "veracrypt -t -d --slot=14"
var ErrNoSuchVolumeMounted = errors.New("no such volume is mounted")

// ErrAdministratorPrivileges gets returned if VeraCrypt fails to obtain administrator privileges, e.g. when trying to
// mount a volume. A program using this library must be started with appropriate privileges to run VeraCrypt
var ErrAdministratorPrivileges = errors.New("failed to obtain administrator privileges")

// ErrNoSuchFileOrDirectory gets returned when trying to mount a volume, but VeraCrypt is unable to locate it
var ErrNoSuchFileOrDirectory = errors.New("no such file or directory")

// ErrNoVolumePath gets returned when trying to call a *Volume func without providing a valid volume path, e.g. an
// empty string. This error is no VeraCrypt error
var ErrNoVolumePath = errors.New("no volume path provided")

// ErrNoVolumesMounted is not an error in the normal sense, since it only indicates that no volumes were mounted.
// This error gets returned when trying to use -l / --list
var ErrNoVolumesMounted = errors.New("no volumes mounted")

// ErrUnknown is returned if an operation failed, but the exact reason cannot be determined. This error is no
// VeraCrypt error
var ErrUnknown = errors.New("an unknown error has occurred")

// ErrCommandNotFound is returned if VeraCrypt (or any other program) is not installed on the operating system (linux).
// This error may differ between systems. This error is no VeraCrypt error
var ErrCommandNotFound = errors.New("command not found")
