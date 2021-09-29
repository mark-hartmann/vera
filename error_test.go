package vera

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrorIs(t *testing.T) {
	err := newError(ErrNoSuchVolumeMounted, "#13")
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrNoSuchVolumeMounted)
	assert.Equal(t, err.Error(), "#13")
	assert.Equal(t, err.Unwrap(), ErrNoSuchVolumeMounted)

	err = newError(ErrNoVolumePath, "no volume found")
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrNoVolumePath)
	assert.Equal(t, err.Error(), "no volume found")
	assert.Equal(t, err.Unwrap(), ErrNoVolumePath)
}
