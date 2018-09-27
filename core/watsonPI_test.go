package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdateProfileWIthContent(t *testing.T) {

	pi, err := NewPersonalityInsight()
	assert.Nil(t, err)

	err = pi.updateProfileWithContent("../resources/test/contents.json")
	assert.Nil(t, err)

	expected := "Conscientiousness"
	value := pi.GetConscientiousness().Name

	assert.Equal(t, expected, value)
}

func TestSaveAndLoadProfileAsJson(t *testing.T) {

	path := "../resources/profile.json"

	pi, err := NewPersonalityInsight()
	assert.Nil(t, err)

	err = pi.updateProfileWithContent("../resources/test/contents.json")
	assert.Nil(t, err)

	err = pi.SaveProfileAsJson(path)
	assert.Nil(t, err)

	err = pi.LoadJsonAsProfile(path)
	assert.Nil(t, err)

	expected := pi.GetAggreeableness().Percentage

	assert.True(t, expected > 0.1)

}
