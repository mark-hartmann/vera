package vera

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// todo: find a way to test this properly
func TestInstalled(t *testing.T) {
	version, ok := Installed()

	assert.True(t, ok)
	assert.NotEmpty(t, version)
	assert.Contains(t, version, "VeraCrypt")
}