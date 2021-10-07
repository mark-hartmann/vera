package vera

import (
	"strconv"
	"strings"
)

const slotMin = 1
const slotMax = 64

// Installed returns whether VeraCrypt is installed or not. If it is not installed, ErrCommandNotFound gets returned.
// If Veracrypt is installed, the version is returned, e.g. VeraCrypt 1.24-Update7
func Installed() (string, error) {

	stdout, err := ExecCommand(version)
	if err != nil {
		return "", err
	}

	// the stdout content has a newline, so we simply trim it away
	return strings.Trim(stdout.String(), "\n"), nil
}

// slotValid returns an ErrParameterIncorrect error if the supplied slot is not valid (at least slotMin and at
// most slotMax). Using this func can prevent an unnecessary cli call
func slotValid(slot uint8) error {
	if slot < slotMin || slot > slotMax {
		msg := ErrParameterIncorrect.Error() + ": " + strconv.Itoa(int(slot))
		return Error{cause: ErrParameterIncorrect, message: msg}
	}

	return nil
}
