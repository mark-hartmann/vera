package vera

import (
	"bytes"
	"os/exec"
	"strings"
)

// newCommand returns a "pre-configured" exec.Cmd plus the stdout and stderr buffers used to intercept
// command line errors and output
func newCommand(args ...Param) (cmd *exec.Cmd, stdout, stderr *bytes.Buffer) {
	stdout = &bytes.Buffer{}
	stderr = &bytes.Buffer{}

	// create new command with
	cmd = exec.Command("veracrypt", genArgs(append([]Param{textOnly, nonInteractive}, args...))...)
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	return
}

// ExecCommandWithStdin adds string to the stdin in of a command, executes the command and returns a buffer with the console output and eventually an
// error, if an error was encountered.
func ExecCommandWithStdin(stdin string, args ...Param) (stdout *bytes.Buffer, err error) {

	cmd, stdout, stderr := newCommand(args...)
	cmd.Stdin = strings.NewReader(stdin)

	if err = cmd.Run(); err != nil {
		return stdout, parseError(stderr.String())
	}

	return stdout, nil
}

// ExecCommand executes a simple command and returns a buffer with the console output and eventually an
// error, if an error was encountered.
func ExecCommand(args ...Param) (stdout *bytes.Buffer, err error) {
	cmd, stdout, stderr := newCommand(args...)

	if err = cmd.Run(); err != nil {
		return stdout, parseError(stderr.String())
	}

	return stdout, nil
}

// genArgs takes a slice of Param structs and generates a slice of strings to pass to the cmd instance
func genArgs(p []Param) []string {
	// add the two required parameters
	var args []string
	var params []string

	for _, param := range p {
		if param.Name == trueCrypt.Name || param.Name == "tc" {
			// prepend --truecrypt to avoid any issues, see "veracrypt --help" for more details
			params = append([]string{param.String()}, params...)
			continue
		}

		// add cli arguments to separate list
		if param.Name == "" {
			args = append(args, param.Value)
		} else {
			str := param.String()
			if len(param.Name) < 3 && strings.Contains(str, " ") {
				params = append(params, strings.Split(str, " ")...)
			} else {
				params = append(params, str)
			}
		}
	}

	// create a new list with the flags and parameters at the beginning and the args at the end
	return append(params, args...)
}
