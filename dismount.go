package vera

// DismountAll dismounts all mounted volumes
func DismountAll() error {
	cmd, _, _ := newCommand(dismount)
	if err := cmd.Run(); err != nil {
		// todo right now it just assumes this is the correct error, create a parseError function later
		return err
	}

	return nil
}

// DismountVolume dismounts a volume using the volume path. To dismount all currently mounted volumes, use DismountAll.
// Does not allow empty strings
func DismountVolume(path string) error {
	if len(path) == 0 {
		return ErrNoVolumePath
	}

	cmd, _, _ := newCommand(dismount, arg(path))
	if err := cmd.Run(); err != nil {
		// todo right now it just assumes this is the correct error, create a parseError function later
		return ErrNoSuchVolumeMounted
	}

	return nil
}
