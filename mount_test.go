package vera

import (
	"os"
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

func (suite *MountTestSuite) TestSlotVolumeMount() {
	props, err := MountSlot("./testdata/basic.vc", 2, "123456789")

	suite.NoError(err)
	suite.NotEqual(MountProperties{}, props)
	suite.Equal(uint8(2), props.Slot)
}

func (suite *MountTestSuite) TestSlotVolumeMountIncorrectPassword() {
	props, err := MountSlot("./testdata/basic.vc", 2, "1234567890") // password is 123456789

	suite.Error(err)
	suite.ErrorIs(err, ErrOperationFailed)
	suite.Equal(MountProperties{}, props)
}

func (suite *MountTestSuite) TestSlotVolumeMountComplexPassword() {
	props, err := MountSlot("./testdata/basic-complex-pw.vc", 2, `s8&"f^T$r'`)

	suite.NoError(err)
	suite.NotEqual(MountProperties{}, props)
	suite.Equal(uint8(2), props.Slot)
}

func (suite *MountTestSuite) TestPathVolumeMount() {
	props, err := MountPath("./testdata/basic.vc", "./testdata/mount", "123456789")
	wd, _ := os.Getwd()

	suite.NoError(err)
	suite.NotEqual(MountProperties{}, props)
	suite.Equal(wd+"/testdata/mount", props.MountPoint)
}

func (suite *MountTestSuite) TestPathVolumeMountpointDoesNotExist() {
	props, err := MountPath("./testdata/basic.vc", "./testdata/mount-non-exist", "123456789")

	suite.Error(err)
	suite.ErrorIs(err, ErrMountPointDoesNotExist)
	suite.Equal(MountProperties{}, props)
}

func (suite *MountTestSuite) TestPathVolumeMountpointNotADirectory() {
	props, err := MountPath("./testdata/basic.vc", "./testdata/testmountpoint", "123456789")

	suite.Error(err)
	suite.ErrorIs(err, ErrMountPointIsNotADirectory)
	suite.Equal(MountProperties{}, props)
}

func (suite *MountTestSuite) TestPathVolumeMountpointInUse() {
	props, err := MountPath("./testdata/basic.vc", "./testdata/mount", "123456789")
	wd, _ := os.Getwd()

	suite.NoError(err)
	suite.NotEqual(MountProperties{}, props)
	suite.Equal(wd+"/testdata/mount", props.MountPoint)

	props, err = MountPath("./testdata/basic2.vc", "./testdata/mount", "123456789")
	suite.Error(err)
	suite.ErrorIs(err, ErrMountPointIsAlreadyInUse)
	suite.Equal(MountProperties{}, props)

}

func (suite *MountTestSuite) TestPathVolumeMountIncorrectPassword() {
	props, err := MountPath("./testdata/basic.vc", "./testdata/mount", "1234567890")

	suite.Error(err)
	suite.ErrorIs(err, ErrOperationFailed)
	suite.Equal(MountProperties{}, props)
}

func (suite *MountTestSuite) TestPathVolumeMountComplexPassword() {
	props, err := MountPath("./testdata/basic-complex-pw.vc", "./testdata/mount", `s8&"f^T$r'`)
	wd, _ := os.Getwd()

	suite.NoError(err)
	suite.NotEqual(MountProperties{}, props)
	suite.Equal(wd+"/testdata/mount", props.MountPoint)
}
