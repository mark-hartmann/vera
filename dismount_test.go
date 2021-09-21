package vera

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type DismountTestSuite struct {
	suite.Suite
}

func TestDismountTestSuite(t *testing.T) {
	suite.Run(t, new(DismountTestSuite))
}

func (suite *DismountTestSuite) BeforeTest(_, _ string) {
	dismountAll()
}

func (suite *DismountTestSuite) AfterTest(_, _ string) {
	dismountAll()
}

// make sure no error is returned if no volume is mounted
func (suite *DismountTestSuite) TestDismountAllDoesNotReturnErrorIfNoVolume() {
	err := DismountAll()
	suite.NoError(err)
}

func (suite *DismountTestSuite) TestDismountVolumeReturnErrorIfEmptyString() {
	err := DismountVolume("")
	suite.Error(err)
	suite.Equal("no volume path provided", err.Error())
}

func (suite *DismountTestSuite) TestDismountVolumeReturnErrorIfUnknownVolume() {
	err := DismountVolume("test.container")
	suite.Error(err)
	suite.ErrorIs(err, ErrNoSuchVolumeMounted)
}

func (suite *DismountTestSuite) TestDismountSlotReturnsErrParameterIncorrectIfInvalidSlot() {
	err := DismountSlot(0)
	suite.Error(err)
	suite.ErrorIs(err, ErrParameterIncorrect)

	err = DismountSlot(65)
	suite.Error(err)
	suite.ErrorIs(err, ErrParameterIncorrect)
}

func (suite *DismountTestSuite) TestDismountSlotReturnsNoErrMountedVolumeDismounted() {
	_, err := Mount("./testdata/basic.vc", 1, Param{Name: "password", Value: "123456789"})
	suite.NoError(err)

	err = DismountSlot(1)
	suite.NoError(err)
}

func (suite *DismountTestSuite) TestDismountVolumeReturnsNoErrMountedVolumeDismounted() {
	_, err := Mount("./testdata/basic.vc", 1, Param{Name: "password", Value: "123456789"})
	suite.NoError(err)

	err = DismountVolume("./testdata/basic.vc")
	suite.NoError(err)
}

func (suite *DismountTestSuite) TestDismountVolumeReturnsErrMountedVolumeDismountedOnlyName() {
	_, err := Mount("./testdata/basic.vc", 1, Param{Name: "password", Value: "123456789"})
	suite.NoError(err)

	err = DismountVolume("basic.vc")
	suite.Error(err)
	suite.ErrorIs(err, ErrNoSuchVolumeMounted)
}

func (suite *DismountTestSuite) TestDismountAllReturnsNoErrMountedVolumeDismounted() {
	_, err := Mount("./testdata/basic.vc", 1, Param{Name: "password", Value: "123456789"})
	suite.NoError(err)

	err = DismountAll()
	suite.NoError(err)
}