package vera

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type CreateTestSuite struct {
	suite.Suite
}

var volumePath = "./testdata/test.vc"

func TestCreateTestSuite(t *testing.T) {
	suite.Run(t, new(CreateTestSuite))
}

func (suite *CreateTestSuite) BeforeTest(_, _ string) {
	DismountAll()
}

func (suite *CreateTestSuite) AfterTest(_, _ string) {
	DismountAll()
	e := os.Remove(volumePath)
	if e != nil {
		fmt.Println("Could not remove testvolume")
		os.Exit(1)
	}
}

func (suite *CreateTestSuite) TestCreate() {
	err := Create(
		volumePath,
		"aes",
		"sha512",
		"ext4",
		"20m",
		"123456789",
	)

	suite.NoError(err)
}

func (suite *CreateTestSuite) TestCreateExistingVolume() {
	err := Create(
		volumePath,
		"aes",
		"sha512",
		"ext4",
		"-32gh",
		"123456789",
	)

	err2 := Create(
		volumePath,
		"aes",
		"sha512",
		"ext4",
		"10m",
		"123456789",
	)

	err3 := Create(
		volumePath,
		"aes",
		"sha512",
		"ext4",
		"10m",
		"123456789",
	)

	suite.ErrorIs(err, ErrParameterIncorrect)
	suite.NoError(err2)
	suite.ErrorIs(err3, ErrVolumePathAlreadyExists)
}
