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

// ErrNoVolumesMounted is not an error in the normal sense, since it only indicates that no volumes were mounted.
// This error gets returned when trying to use -l / --list
var ErrNoVolumesMounted = errors.New("no volumes mounted")

// ErrOperationFailed is returned when VeraCrypt failed to mount a volume due to faulty details, such as incorrect
// password or passing a file path that is not a volume. Unfortunately VeraCrypt does not tell the exact cause of this
// error, so it is up to the user to check the details they used
var ErrOperationFailed = errors.New("operation failed")

// ErrUnknown is returned if an operation failed, but the exact reason cannot be determined. This error is no
// VeraCrypt error
var ErrUnknown = errors.New("an unknown error has occurred")

// ErrCommandNotFound is returned if VeraCrypt (or any other program) is not installed on the operating system (linux).
// This error may differ between systems. This error is no VeraCrypt error
var ErrCommandNotFound = errors.New("command not found")

// ErrNoVolumePath gets returned when trying to call a *Volume func without providing a valid volume path, e.g. an
// empty string. This error is no VeraCrypt error
var ErrNoVolumePath = errors.New("no volume path provided")

// ErrMountPointDoesNotExist is returned when mounting a volume to a mount point that does not exist
var ErrMountPointDoesNotExist = errors.New("mount point does not exist")

// ErrMountPointIsNotADirectory is returned when the mount point is not a directory
var ErrMountPointIsNotADirectory = errors.New("mount point is not a directory")

// ErrMountPointIsAlreadyInUse is returned when the mount point is already mounted by another volume
var ErrMountPointIsAlreadyInUse = errors.New("mount point is already in use")

// ErrVolumeAlreadyMounted is returned when the volume is already mounted
var ErrVolumeAlreadyMounted = errors.New("volume already mounted")

// ErrVolumePathAlreadyExists gets returned when creating a volume and the path for the volume already exists.
// This is not a Veracrypt error.
var ErrVolumePathAlreadyExists = errors.New("the path for creating the volume already exists")
