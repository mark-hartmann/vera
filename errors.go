package vera

import "errors"

var ErrNoVolumesMounted = errors.New("no volumes mounted")
var ErrNoSuchVolumeMounted = errors.New("no such value is mounted")
