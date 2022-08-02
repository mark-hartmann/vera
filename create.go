package vera

import "os"

// Create creates a new Veracrypt volume
// volumePath is the full path to a volume, including the volume name
func Create(volumePath string, cipher string, hash string, filesystem string, size string, password string, opts ...Param) error {

	// Veracrypt happily overwrites existing files or directories, we prevent that from happening
	if _, err := os.Stat(volumePath); !os.IsNotExist(err) {
		return ErrVolumePathAlreadyExists
	}

	//Create a separate slice to make sure create is the first opt
	var createOpts []Param
	createOpts = append(
		createOpts,
		Param{Name: "create", IsFlag: true},
		Param{Value: volumePath},
		Param{Name: "encryption", Value: cipher},
		Param{Name: "hash", Value: hash},
		Param{Name: "filesystem", Value: filesystem},
		Param{Name: "size", Value: size},
		Param{Name: "stdin", IsFlag: true},
	)

	opts = append(createOpts, opts...)

	_, err := ExecCommandWithStdin(password, opts...)

	if err != nil {
		return err
	}
	return nil
}
