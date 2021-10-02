package vera

import (
	"github.com/stretchr/testify/assert"
	"testing"
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