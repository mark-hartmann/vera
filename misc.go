package vera

import (
	"bytes"
	"os/exec"
	"strings"
)

// Installed returns whether VeraCrypt is installed or not, as well es the installed
// version, e.g VeraCrypt 1.24-Update7
func Installed() (string, bool) {
	cmd, stdout, _ := newCommand("-t", "--version")
	// cmd.Run will return "veracrypt: command not found" if veracrypt is not
	// installed on this system
	if err := cmd.Run(); err != nil {
		return "", false
	}

	// the stdout content has a newline, so we simply trim it away
	return strings.Trim(stdout.String(), "\n"), true
}

// newCommand returns a "pre configured" exec.Cmd plus the stdout and stderr buffers
// used to intercept command line errors and output
func newCommand(args ...string) (cmd *exec.Cmd, stdout, stderr *bytes.Buffer) {
	stdout = &bytes.Buffer{}
	stderr = &bytes.Buffer{}

	cmd = exec.Command("veracrypt", args...)
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	return
}
