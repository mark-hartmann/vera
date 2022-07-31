package vera

import (
	"path"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ListTestSuite struct {
	suite.Suite
}

func TestListTestSuite(t *testing.T) {
	suite.Run(t, new(ListTestSuite))
}

// dismount all mounted volumes
func (suite *ListTestSuite) BeforeTest(_, _ string) {
	DismountAll()
}

// dismount all mounted volumes
func (suite *ListTestSuite) AfterTest(_, _ string) {
	DismountAll()
}

// make sure list returns an empty array even if VeraCrypt returns an ErrNoVolumesMounted error. If no volumes are
// mounted, the returned error must be of type ErrNoVolumesMounted
func (suite ListTestSuite) TestListReturnsEmptySliceIfNoMount() {
	mounts, err := List()

	suite.Error(err)
	suite.ErrorIs(err, ErrNoVolumesMounted)
	suite.NotNil(mounts)
	suite.Empty(mounts)
}

// make sure ErrNoSuchVolumeMounted is returned of no volumes are mounted
func (suite ListTestSuite) TestPropertiesSlotErrNoSuchVolumeMounted() {
	_, err := PropertiesSlot(1)
	suite.ErrorIs(err, ErrNoSuchVolumeMounted)

	_, err = PropertiesSlot(64)
	suite.ErrorIs(err, ErrNoSuchVolumeMounted)
}

// make sure ErrParameterIncorrect is returned if the slot number is less than one. The VeraCrypt slot range starts
// at one
func (suite ListTestSuite) TestPropertiesSlotDoesNotAcceptValuesBelowOne() {
	_, err := PropertiesSlot(0)
	suite.ErrorIs(err, ErrParameterIncorrect)
}

// make sure that the VeraCrypt slot range is well respected. If no volumes are mounted, and we try to retrieve the
// properties of slot one, ErrNoSuchVolumeMounted must be returned. If the slot is 65 and out of range, we
// expect ErrParameterIncorrect
func (suite ListTestSuite) TestPropertiesSlotSlotOutOfBoundsErrParameterIncorrect() {
	// slot 64 is not out of bounds, so we expect an ErrNoSuchVolumeMounted error
	_, err := PropertiesSlot(64)
	suite.ErrorIs(err, ErrNoSuchVolumeMounted)

	// VeraCrypt does not support more than 64 mounted volumes
	_, err = PropertiesSlot(65)
	suite.ErrorIs(err, ErrParameterIncorrect)
}

// check if we get the correct errors while calling the PropertiesVolume func without any mounted volumes
func (suite ListTestSuite) TestPropertiesVolumeErrNoSuchVolumeMounted() {
	const volume = "./testdata/basic.vc"
	_, err := PropertiesVolume(volume)
	suite.ErrorIs(err, ErrNoSuchVolumeMounted)
	_, err = PropertiesVolume(volume)
	suite.ErrorIs(err, ErrNoSuchVolumeMounted)
}

// check if the PropertiesSlot func returns a MountProperties struct that matches the mounted volume
func (suite ListTestSuite) TestPropertiesSlotReturnsCorrectMountProperties() {
	const volume = "./testdata/basic.vc"
	mountProps, err := Mount(volume, 1, "123456789")
	suite.NoError(err)

	props, err := PropertiesSlot(1)
	suite.NoError(err)
	suite.Equal(mountProps, props)
	// use path.Clean(volume) to remove the relative dot
	suite.Contains(props.Volume, path.Clean(volume))
}

// check if the PropertiesVolume func returns a MountProperties struct that matches the mounted volume
func (suite ListTestSuite) TestPropertiesVolumeReturnsCorrectMountProperties() {
	const volume = "./testdata/basic.vc"
	mountProps, err := Mount(volume, 1, "123456789")
	suite.NoError(err)

	props, err := PropertiesVolume(volume)
	suite.NoError(err)
	suite.Equal(mountProps, props)
	// use path.Clean(volume) to remove the relative dot
	suite.Contains(props.Volume, path.Clean(volume))
}

// make sure both PropertiesSlot and PropertiesVolume are returning the same data
func (suite ListTestSuite) TestPropertiesVolumeAndPropertiesSlotReturnTheSameData() {
	const volume = "./testdata/basic.vc"
	_, err := Mount(volume, 1, "123456789")
	suite.NoError(err)

	propsSlot, _ := PropertiesSlot(1)
	propsVolume, _ := PropertiesVolume(path.Clean(volume))
	suite.Equal(propsSlot, propsVolume)
}
