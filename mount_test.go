package vera

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type MountTestSuite struct {
	suite.Suite
}

func TestMountTestSuite(t *testing.T) {
	suite.Run(t, new(MountTestSuite))
}

func (suite *MountTestSuite) BeforeTest(suiteName, testName string) {
	dismountAll()
}

func (suite *MountTestSuite) AfterTest(suiteName, testName string) {
	dismountAll()
}

func (suite *MountTestSuite) TestBasicContainerMount() {
	_, err := Mount("./testdata/basic.vc", 2, Param{Name: "password", Value: "123456789"})
	suite.NoError(err)

	props, err := List()

	suite.NoError(err)
	suite.NotEmpty(props)
	suite.Len(props, 1)
}

// dismountAll dismounts all currently mounted volumes
func dismountAll() {
	DismountAll()
}
