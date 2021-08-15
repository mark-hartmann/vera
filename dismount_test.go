package vera

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDismountAllDoesNotReturnErrorIfNoVolume(t *testing.T) {
	err := DismountAll()
	assert.NoError(t, err)
}

func TestDismountVolumeReturnErrorIfEmptyString(t *testing.T) {
	err := DismountVolume("")
	assert.Error(t, err)
	assert.Equal(t, "no volume path provided", err.Error())
}

func TestDismountVolumeReturnErrorIfUnknownVolume(t *testing.T) {
	err := DismountVolume("test.container")
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrNoSuchVolumeMounted)
}

func TestDismountSlotReturnsErrParameterIncorrectIfInvalidSlot(t *testing.T) {
	err := DismountSlot(0)
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrParameterIncorrect)

	err = DismountSlot(65)
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrParameterIncorrect)
}