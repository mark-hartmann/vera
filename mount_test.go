package vera

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type MountTestSuite struct {
	suite.Suite
}

func TestMountTestSuite(t *testing.T) {
	suite.Run(t, new(MountTestSuite))
}

func (suite *MountTestSuite) BeforeTest(suiteName, testName string) {
	DismountAll()
}

func (suite *MountTestSuite) AfterTest(suiteName, testName string) {
	DismountAll()
}

func (suite MountTestSuite) TestSlotOutOfBoundsErrParameterIncorrect() {
	emptyProps := MountProperties{}

	// there is no slot 0
	props, err := MountSlot("./testdata/basic.vc", 0, "123456789")
	suite.Equal(emptyProps, props)
	suite.ErrorIs(err, ErrParameterIncorrect)

	// VeraCrypt only supports 64 slots
	_, err = MountSlot("./testdata/basic.vc", 65, "123456789")
	suite.Equal(emptyProps, props)
	suite.ErrorIs(err, ErrParameterIncorrect)
}

func (suite *MountTestSuite) TestBasicVolumeMount() {
	props, err := MountSlot("./testdata/basic.vc", 2, "123456789")

	suite.NoError(err)
	suite.NotEqual(MountProperties{}, props)
	suite.Equal(uint8(2), props.Slot)
}

func (suite *MountTestSuite) TestBasicVolumeMountIncorrectPassword() {
	props, err := MountSlot("./testdata/basic.vc", 2, "1234567890") // password is 123456789

	suite.Error(err)
	suite.ErrorIs(err, ErrOperationFailed)
	suite.Equal(MountProperties{}, props)
}

func (suite *MountTestSuite) TestBasicVolumeMountComplexPassword() {
	props, err := MountSlot("./testdata/basic-complex-pw.vc", 2, `s8&"f^T$r'`)

	suite.NoError(err)
	suite.NotEqual(MountProperties{}, props)
	suite.Equal(uint8(2), props.Slot)
}
