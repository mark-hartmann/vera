package vera

import "errors"

var ErrNoVolumePath = errors.New("no volume path provided")
var ErrNoVolumesMounted = errors.New("no volumes mounted")
var ErrNoSuchVolumeMounted = errors.New("no such value is mounted")
