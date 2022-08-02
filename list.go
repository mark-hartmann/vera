package vera

import (
	"regexp"
	"strconv"
	"strings"
)

// MountProperties holds basic information about a mounted VeraCrypt volume
type MountProperties struct {
	Slot       uint8
	Volume     string
	MountPoint string
}

// List returns a list of all mounted volumes. If no value is mounted an empty list is returned, as well
// as a ErrNoVolumesMounted error
func List() ([]MountProperties, error) {

	stdout, err := ExecCommand(list)
	if err != nil {
		return make([]MountProperties, 0), err
	}

	return parseListOutput(stdout.String()), nil
}

// PropertiesSlot returns a MountProperties struct for the volume mounted in the given slot. This function will
// return an error if the slot is empty or out of bounds (1-64)
func PropertiesSlot(slot uint8) (MountProperties, error) {
	if err := slotValid(slot); err != nil {
		return MountProperties{}, err
	}

	stdout, err := ExecCommand(list, Param{Name: "slot", Value: strconv.Itoa(int(slot))})
	if err != nil {
		return MountProperties{}, err
	}

	// parse the stdout, because we used the slot flag, only one entry is returned
	return parseListOutput(stdout.String())[0], nil
}

// PropertiesMountpoint returns a MountProperties struct for the volume mounted on the given mountPoint
func PropertiesMountpoint(mountPoint string) (MountProperties, error) {

	stdout, err := ExecCommand(list, Param{Value: mountPoint})
	if err != nil {
		return MountProperties{}, err
	}

	// parse the stdout, because we used the slot flag, only one entry is returned
	return parseListOutput(stdout.String())[0], nil
}

// PropertiesVolume returns a MountProperties struct for the volume mounted in the given slot.
func PropertiesVolume(volume string) (MountProperties, error) {
	if len(volume) == 0 {
		return MountProperties{}, ErrNoVolumePath
	}

	stdout, err := ExecCommand(list, Param{Value: volume})
	if err != nil {
		return MountProperties{}, err
	}

	// parse the stdout, because we used the volume param, only one entry is returned
	return parseListOutput(stdout.String())[0], nil
}

// parseListOutput takes the content of the commands stdOut and parses it
func parseListOutput(s string) []MountProperties {
	regex := regexp.MustCompile(`(?m)^(?P<slot>\d): (?P<volume>.*) (?P<vDevice>.*) (?P<mount>.*)$`)
	var result []MountProperties

	for _, entry := range strings.SplitN(s, "\n", strings.Count(s, "\n")) {
		mapped := generateCaptureMap(entry, regex)
		slot, err := strconv.ParseInt(mapped["slot"], 10, 8)
		if err != nil {
			return result
		}

		result = append(result, MountProperties{
			Slot:       uint8(slot),
			Volume:     mapped["volume"],
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
