package vera

import (
	"bytes"
	"errors"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

// MountProperties holds basic information about a mounted VeraCrypt container
type MountProperties struct {
	Slot       uint8
	Container  string
	MountPoint string
}

func List() ([]MountProperties, error) {
	cmd := exec.Command("veracrypt", "-t", "-l")
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		// parse error from stderr and return this instead of the err itself, which is "exit status 1"
		return []MountProperties{}, errors.New(stderr.String())
	}

	return parseListOutput(stdout.String()), nil
}

// parseListOutput takes the content of the commands stdOut and parses it
func parseListOutput(s string) []MountProperties {
	regex := regexp.MustCompile(`(?m)^(?P<slot>\d): (?P<container>.*) (?P<vDevice>.*) (?P<mount>.*)$`)
	var result []MountProperties

	for _, entry := range strings.SplitN(s, "\n", strings.Count(s, "\n")) {
		mapped := generateCaptureMap(entry, regex)
		slot, err := strconv.ParseInt(mapped["slot"], 10, 8)
		if err != nil {
			return result
		}

		result = append(result, MountProperties{
			Slot:       uint8(slot),
			Container:  mapped["container"],
			MountPoint: mapped["mount"],
		})
	}

	return result
}

// generateCaptureMap generates a map of matches using the capture groups contained in the given regexp.Regexp
func generateCaptureMap(s string, regex *regexp.Regexp) map[string]string {
	results := map[string]string{}
	matches := regex.FindStringSubmatch(s)

	if len(matches) == 0 {
		return results
	}

	for i, name := range regex.SubexpNames() {
		if i != 0 && name != "" {
			results[name] = matches[i]
		}
	}

	return results
}
