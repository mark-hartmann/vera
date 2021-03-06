package vera

import "strconv"

// DismountAll dismounts all mounted volumes
func DismountAll() error {
	_, err := ExecCommand(dismount)
	return err
}

// DismountSlot dismounts a volume using it's assigned mount slot. The mount slot range is 1-64
func DismountSlot(slot uint8) error {
	if err := slotValid(slot); err != nil {
		return err
	}

	_, err := ExecCommand(dismount, Param{Name: "slot", Value: strconv.Itoa(int(slot))})
	return err
}

// DismountVolume dismounts a volume using the volume path. To dismount all currently mounted volumes, use DismountAll.
// Does not allow empty strings
func DismountVolume(path string) error {
	if len(path) == 0 {
		return ErrNoVolumePath
	}

	_, err := ExecCommand(dismount, arg(path))
	return err
}
