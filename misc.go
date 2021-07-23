package vera

import (
	"bytes"
	"os/exec"
	"strings"
)

// Installed returns whether VeraCrypt is installed or not, as well es the installed
// version, e.g VeraCrypt 1.24-Update7
func Installed() (string, bool) {
	cmd := exec.Command("veracrypt", "-t", "--version")
	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	if err := cmd.Run(); err != nil {
		return "", false
	}

	return strings.Trim(stdout.String(), "\n"), true
}