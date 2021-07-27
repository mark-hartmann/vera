package vera

import "errors"

// DismountAll dismounts all mounted volumes
func DismountAll() error {
	cmd, _, _ := newCommand("-d")
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

// DismountVolume dismounts a volume using the volume path. Does not allow empty strings, to dismount all currently
// mounted volumes, use DismountAll
func DismountVolume(path string) error {
	if len(path) == 0 {
		return errors.New("no volume path provided")
	}
	cmd, _, _ := newCommand("-d", path)
	if err := cmd.Run(); err != nil {
		// todo right now it just assumes this is the correct error, create a parseError function later
		return ErrNoSuchVolumeMounted
	}

	return nil
}
