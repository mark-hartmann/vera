package vera

import (
	"strconv"
)

// MountSlot mounts the volume in the given slot. Depending on the volume, the appropriate parameters must be set.
// password is passed through stdin
func MountSlot(volume string, slot uint8, password string, opts ...Param) (MountProperties, error) {
	// check if the slot number is supported
	if err := slotValid(slot); err != nil {
		return MountProperties{}, err
	}

	// append path and mount as Param
	opts = append(opts, Param{Value: volume}, Param{Name: "slot", Value: strconv.Itoa(int(slot))}, Param{Name: "stdin", IsFlag: true})
	_, err := ExecCommandWithStdin(password, opts...)
	if err != nil {
		return MountProperties{}, err
	}

	return PropertiesVolume(volume)
}

// MountPath mounts the volume in the given path. Depending on the volume, the appropriate parameters must be set.
// password is passed through stdin
func MountPath(volume string, path string, password string, opts ...Param) (MountProperties, error) {

	// append path and mount as Param
	opts = append(opts, Param{Value: volume}, Param{Value: path}, Param{Name: "stdin", IsFlag: true})
	_, err := ExecCommandWithStdin(password, opts...)
	if err != nil {
		return MountProperties{}, err
	}

	return PropertiesVolume(volume)
}
