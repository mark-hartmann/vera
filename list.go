package vera

import (
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

// List returns a list of all mounted volumes. If no value is mounted a empty list is returned, as well
// as a ErrNoVolumesMounted error
func List() ([]MountProperties, error) {
	cmd, stdout, _ := newCommand(list)
	if err := cmd.Run(); err != nil {
		// parse error from stderr and return this instead of the error itself, which is "exit status 1"
		return make([]MountProperties, 0), ErrNoVolumesMounted
	}

	return parseListOutput(stdout.String()), nil
}

// PropertiesSlot returns a MountProperties struct for the container mounted in the given slot. This function will
// return an error if the slot is empty or out of bounds (max 64)
func PropertiesSlot(slot uint8) (MountProperties, error) {
	if slot < 1 || slot > 64 {
		return MountProperties{}, ErrParameterIncorrect
	}

	cmd, stdout, _ := newCommand(list, Param{Name: "slot", Value: strconv.Itoa(int(slot))})
	if err := cmd.Run(); err != nil {
		// we handled the valid slot range above and --list does not require su privileges, so we can simply return
		// ErrNoSuchVolumeMounted
		return MountProperties{}, ErrNoSuchVolumeMounted
	}

	// parse the stdout, because we used the slot flag, only one entry is returned
	return parseListOutput(stdout.String())[0], nil
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
