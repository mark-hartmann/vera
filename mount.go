package vera

import (
	"strconv"
)

// Mount mounts the volume in the given slot. Depending on the volume or container, the appropriate parameters
// must be set, but at least the password
func Mount(container string, slot uint8, opts ...Param) (MountProperties, error) {
	// check if the slot number is supported
	if slot < SlotMin || slot > SlotMax {
		return MountProperties{}, ErrParameterIncorrect
	}

	// append path and mount as Param
	opts = append(opts, Param{Value: container}, Param{Name: "slot", Value: strconv.Itoa(int(slot))})
	cmd, _, _ := newCommand(opts...)

	if err := cmd.Run(); err != nil {
		return MountProperties{}, err
	}

	return PropertiesSlot(slot)
}
