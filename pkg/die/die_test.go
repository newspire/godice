package die

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DieTestSuite struct {
	suite.Suite
}

func (s *DieTestSuite) SetupTest() {
	rand.Seed(42)
}

func (s *DieTestSuite) TestNewDieDefaultsToSixSides() {
	d := NewDie(0)

	assert.False(s.T(), d.IsSet(), "new die should not be marked as set after initialization")
	assert.GreaterOrEqual(s.T(), d.Value(), 1, "default die should roll at least 1")
	assert.LessOrEqual(s.T(), d.Value(), 6, "default die should roll at most 6")
}

func (s *DieTestSuite) TestNewDieUsesCustomSides() {
	d := NewDie(12)

	assert.False(s.T(), d.IsSet(), "new die should not be marked as set after initialization")
	assert.GreaterOrEqual(s.T(), d.Value(), 1, "custom die should roll at least 1")
	assert.LessOrEqual(s.T(), d.Value(), 12, "custom die should roll at most sides")
}

func (s *DieTestSuite) TestRollResetsIsSet() {
	d := NewDie(6)
	d.Set(4)

	assert.True(s.T(), d.IsSet(), "die should be marked as set after Set")

	rolled := d.Roll()

	assert.False(s.T(), d.IsSet(), "die should not be marked as set after Roll")
	assert.GreaterOrEqual(s.T(), rolled, 1, "rolled value should be at least 1")
	assert.LessOrEqual(s.T(), rolled, 6, "rolled value should be at most 6")
	assert.Equal(s.T(), rolled, d.Value(), "Value should return the last rolled value")
}

func (s *DieTestSuite) TestSetStoresValue() {
	d := NewDie(6)

	value := d.Set(3)

	assert.True(s.T(), d.IsSet(), "die should be marked as set after Set")
	assert.Equal(s.T(), 3, value, "Set should return the stored value")
	assert.Equal(s.T(), 3, d.Value(), "Value should return the value set by Set")
}

func TestDieTestSuite(t *testing.T) {
	suite.Run(t, new(DieTestSuite))
}
