package vera

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// todo: find a way to test this properly
func TestInstalled(t *testing.T) {
	version, err := Installed()

	assert.NoError(t, err)
	assert.NotEmpty(t, version)
	assert.Contains(t, version, "VeraCrypt")
}

func TestSlotValid(t *testing.T) {

	err := slotValid(SlotMin)
	assert.NoError(t, err)
	assert.Nil(t, err)

	err = slotValid(SlotMax)
	assert.NoError(t, err)
	assert.Nil(t, err)

	err = slotValid(0)
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrParameterIncorrect)
	assert.Equal(t, "parameter incorrect: 0", err.Error())

	err = slotValid(65)
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrParameterIncorrect)
	assert.Equal(t, "parameter incorrect: 65", err.Error())
}