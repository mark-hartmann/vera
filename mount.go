package vera

import (
	"strconv"
)

// Mount mounts the volume in the given slot. Depending on the volume, the appropriate parameters must be set.
// password is passed through stdin
func Mount(volume string, slot uint8, password string, opts ...Param) (MountProperties, error) {
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

	return PropertiesSlot(slot)
}
