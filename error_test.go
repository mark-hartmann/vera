package vera

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorIs(t *testing.T) {
	err := newError(ErrNoSuchVolumeMounted, "#13")
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrNoSuchVolumeMounted)
	assert.Equal(t, err.Error(), "#13")
	assert.Equal(t, err.Unwrap(), ErrNoSuchVolumeMounted)

	err = newError(ErrNoVolumePath, "no volume found")

	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrNoVolumePath)
	assert.Equal(t, err.Error(), "no volume found")
	assert.Equal(t, err.Unwrap(), ErrNoVolumePath)
}

func TestParseErrIncorrectCommandLineSpecified(t *testing.T) {

	errMap := map[string]string{
		`unexpected parameter 'test123'`: `  --use-dummy-sudo-password     Use dummy password in sudo to detect if it is already authenticated
Unexpected parameter 'test123'
Error: Incorrect command line specified.`,
		`unknown option 'x'`: `  --use-dummy-sudo-password     Use dummy password in sudo to detect if it is already authenticated
Unknown option 'x'
Error: Incorrect command line specified.`,
		`unknown long option 'unknown'`: `  --use-dummy-sudo-password     Use dummy password in sudo to detect if it is already authenticated
Unknown long option 'unknown'
Error: Incorrect command line specified.`,
		`option 'm' requires a value`: `  --use-dummy-sudo-password     Use dummy password in sudo to detect if it is already authenticated
Option 'm' requires a value.
Error: Incorrect command line specified.`,
	}

	for msg, stderr := range errMap {
		err := parseError(stderr)

		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrIncorrectCommandLineSpecified)
		assert.Equal(t, msg, err.Error())
		assert.Equal(t, ErrIncorrectCommandLineSpecified.Error(), err.Unwrap().Error())
	}
}

func TestErrUnknownOption(t *testing.T) {

	errMap := map[string]string{
		`unknown option: test`:   `Error: Unknown option: test`,
		`unknown option: test.`:  `Error: Unknown option: test.`,
		`unknown option: #test.`: `Error: Unknown option: #test.`,
	}

	for msg, stderr := range errMap {
		err := parseError(stderr)

		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrUnknownOption)
		assert.Equal(t, msg, err.Error())
		assert.Equal(t, ErrUnknownOption.Error(), err.Unwrap().Error())
	}
}

func TestErrParameterIncorrect(t *testing.T) {

	errMap := map[string]string{
		`parameter incorrect: one`: `Error: Parameter incorrect: one`, // --slot=one
		`parameter incorrect: 65`:  `Error: Parameter incorrect: 65`,  // --slot=65
	}

	for msg, stderr := range errMap {
		err := parseError(stderr)

		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrParameterIncorrect)
		assert.Equal(t, msg, err.Error())
		assert.Equal(t, ErrParameterIncorrect.Error(), err.Unwrap().Error())
	}
}

func TestErrNoSuchVolumeMounted(t *testing.T) {
	err := parseError("Error: No such volume is mounted.")
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrNoSuchVolumeMounted)
	assert.Equal(t, "no such volume is mounted", err.Error())
	assert.Equal(t, ErrNoSuchVolumeMounted.Error(), err.Unwrap().Error())
}

func TestErrAdministratorPrivileges(t *testing.T) {
	err := parseError("Error: Failed to obtain administrator privileges.")
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrAdministratorPrivileges)
	assert.Equal(t, "failed to obtain administrator privileges", err.Error())
	assert.Equal(t, ErrAdministratorPrivileges.Error(), err.Unwrap().Error())
}

func TestErrNoSuchFileOrDirectory(t *testing.T) {

	errMap := map[string]string{
		`no such file or directory: /home/mark/repos/mark-hartmann/vera/test.vc`: `Error: No such file or directory:
/home/mark/repos/mark-hartmann/vera/test.vc

VeraCrypt::File::Open:232`,
		`no such file or directory: /home/mark/Documents/important-stuff.vc`: `Error: No such file or directory:
/home/mark/Documents/important-stuff.vc

VeraCrypt::File::Open:232`,
	}

	for msg, stderr := range errMap {
		err := parseError(stderr)

		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrNoSuchFileOrDirectory)
		assert.Equal(t, msg, err.Error())
		assert.Equal(t, ErrNoSuchFileOrDirectory.Error(), err.Unwrap().Error())
	}
}

func TestErrNoVolumesMounted(t *testing.T) {
	err := parseError("Error: No volumes mounted.")
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrNoVolumesMounted)
	assert.Equal(t, "no volumes mounted", err.Error())
	assert.Equal(t, ErrNoVolumesMounted.Error(), err.Unwrap().Error())
}

func TestErrUnknown(t *testing.T) {

	errMap := []string{
		``,
		`Error: Something went wrong`,
		`some other fatal error`,
	}

	for _, stderr := range errMap {
		err := parseError(stderr)

		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrUnknown)
		assert.Equal(t, ErrUnknown.Error(), err.Error())
		assert.Equal(t, ErrUnknown.Error(), err.Unwrap().Error())
	}
}

func TestErrCommandNotFound(t *testing.T) {
	err := parseError("veracrypt: command not found")
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrCommandNotFound)
	assert.Equal(t, "veracrypt: command not found", err.Error())
	assert.Equal(t, ErrCommandNotFound.Error(), err.Unwrap().Error())
}

func TestErrOperationFailed(t *testing.T) {
	err := parseError(`Error: Operation failed due to one or more of the following:
- Incorrect password.
- Incorrect Volume PIM number.
- Incorrect PRF (hash).
- Not a valid volume.`)
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrOperationFailed)
	assert.Equal(t, "operation failed", err.Error())
	assert.Equal(t, ErrOperationFailed.Error(), err.Unwrap().Error())
}

func TestErrMountPointDoesNotExist(t *testing.T) {
	err := parseError("Error: mount: /my-non-existing/mountpoint: mount point does not exist.")
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrMountPointDoesNotExist)
	assert.Equal(t, "/my-non-existing/mountpoint: mount point does not exist", err.Error())
	assert.Equal(t, ErrMountPointDoesNotExist.Error(), err.Unwrap().Error())
}

func TestErrMountPointIsNotADirectory(t *testing.T) {
	err := parseError("Error: mount: /my-file-to-mount: mount point is not a directory.")
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrMountPointIsNotADirectory)
	assert.Equal(t, "/my-file-to-mount: mount point is not a directory", err.Error())
	assert.Equal(t, ErrMountPointIsNotADirectory.Error(), err.Unwrap().Error())
}

func TestErrMountPointIsAlreadyInUse(t *testing.T) {
	err := parseError("Error: Mount point is already in use.")
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrMountPointIsAlreadyInUse)
	assert.Equal(t, "mount point is already in use", err.Error())
	assert.Equal(t, ErrMountPointIsAlreadyInUse.Error(), err.Unwrap().Error())
}

func TestErrVolumeAlreadyMounted(t *testing.T) {
	err := parseError("Error: The volume you are trying to mount is already mounted.")
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrVolumeAlreadyMounted)
	assert.Equal(t, "the volume you are trying to mount is already mounted", err.Error())
	assert.Equal(t, ErrVolumeAlreadyMounted.Error(), err.Unwrap().Error())
}
