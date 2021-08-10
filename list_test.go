package vera

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// todo: write more tests once the Mount function is written
func TestListReturnsEmptyMapIfNoMount(t *testing.T) {
	mounts, err := List()

	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrNoVolumesMounted)
	assert.NotNil(t, mounts)
	assert.Empty(t, mounts)
}

func TestPropertiesSlotErrNoSuchVolumeMounted(t *testing.T) {
	_, err := PropertiesSlot(1)
	assert.ErrorIs(t, err, ErrNoSuchVolumeMounted)
	_, err = PropertiesSlot(64)
	assert.ErrorIs(t, err, ErrNoSuchVolumeMounted)
}

func TestPropertiesSlotDoesNotAcceptValuesBelowOne(t *testing.T) {
	_, err := PropertiesSlot(0)
	assert.ErrorIs(t, err, ErrParameterIncorrect)
}

func TestPropertiesSlotSlotOutOfBoundsErrParameterIncorrect(t *testing.T) {
	// slot 64 is not out of bounds, so we expect an ErrNoSuchVolumeMounted error
	_, err := PropertiesSlot(64)
	assert.ErrorIs(t, err, ErrNoSuchVolumeMounted)
	// VeraCrypt does not support more than 64 mounted containers
	_, err = PropertiesSlot(65)
	assert.ErrorIs(t, err, ErrParameterIncorrect)
}

func TestPropertiesVolumeErrNoSuchVolumeMounted(t *testing.T) {
	_, err := PropertiesVolume("./testdata/basic.vc")
	assert.ErrorIs(t, err, ErrNoSuchVolumeMounted)
	_, err = PropertiesVolume("./testdata/basic.vc")
	assert.ErrorIs(t, err, ErrNoSuchVolumeMounted)
}