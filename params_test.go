package vera

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParam_String(t *testing.T) {
	testData := []struct {
		param    Param
		expected string
	}{
		{Param{Name: "auto-mount", Value: "foo"}, "--auto-mount=foo"},
		{Param{Name: "backup-headers", IsFlag: true}, "--backup-headers"},
		{Param{Name: "C", IsFlag: true}, "-C"},
		{Param{Name: "c", IsFlag: true}, "-c"},
		{Param{Name: "p", Value: "123456789"}, "-p 123456789"}, // -p is no flag
	}

	for _, data := range testData {
		assert.Equal(t, data.expected, data.param.String())
	}
}
