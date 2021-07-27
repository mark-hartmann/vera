package vera

// DismountAll dismounts all mounted volumes
func DismountAll() error {
	cmd, _, _ := newCommand("-d")
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}