package vera

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// todo: write more tests once the Mount function is written
func TestListReturnsEmptyMapIfNoMount(t *testing.T) {
	mounts := List()

	assert.NotNil(t, mounts)
	assert.Empty(t, mounts)
}