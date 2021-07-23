package vera

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInstalled(t *testing.T) {
	version, ok := Installed()

	assert.True(t, ok)
	assert.NotEmpty(t, version)
	assert.Contains(t, version, "VeraCrypt")
}