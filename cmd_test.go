package vera

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenArgs(t *testing.T) {
	params := []Param{
		{Name: "t", IsFlag: true},
		{Name: "non-interactive", IsFlag: true},
		{Name: "p", Value: "123456789"},
		{Name: "pim", Value: "0"},
	}

	assert.Equal(t, []string{"-t", "--non-interactive", "-p", "123456789", "--pim=0"}, genArgs(params))
}

func TestGenArgsHandlesArguments(t *testing.T) {
	params := []Param{
		{Name: "t", IsFlag: true},
		{Value: "./test"}, // arguments are the last entries in the slice and be in order
		{Name: "pim", Value: "0"},
		{Value: "./test2"},
	}

	assert.Equal(t, []string{"-t", "--pim=0", "./test", "./test2"}, genArgs(params))
}

func TestGenArgsHandlesTrueCryptAsSpecialParam(t *testing.T) {
	params := []Param{
		{Name: "t", IsFlag: true},
		{Name: "non-interactive", IsFlag: true},
		{Name: "p", Value: "123456789"},
		{Name: "tc", IsFlag: true}, // tc is treated special and should be the first entry of the generated slice
	}

	assert.Equal(t, []string{"-tc", "-t", "--non-interactive", "-p", "123456789"}, genArgs(params))
}
