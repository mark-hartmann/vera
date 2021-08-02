package vera

import (
	"strings"
)

// Installed returns whether VeraCrypt is installed or not, as well es the installed
// version, e.g. VeraCrypt 1.24-Update7
func Installed() (string, bool) {
	cmd, stdout, _ := newCommand(version)
	// cmd.Run will return "veracrypt: command not found" if veracrypt is not
	// installed on this system
	if err := cmd.Run(); err != nil {
		return "", false
	}

	// the stdout content has a newline, so we simply trim it away
	return strings.Trim(stdout.String(), "\n"), true
}
