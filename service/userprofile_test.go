package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSaveProfile(t *testing.T) {

	pi, err := NewPersonalityInsight()
	assert.Nil(t, err)

	profile, err := pi.GetUserProfile("../resources/test/contents.json")
	assert.Nil(t, err)

	err = profile.SaveProfileAsJson("../resources/test/SaveProfile.json")
	assert.Nil(t, err)

}

func TestSaveAndLoadProfileAsJson(t *testing.T) {

	path := "../resources/test/profile2.json"

	userProfile := UserProfile{}

	err := userProfile.LoadJsonAsProfile("../resources/test/SaveProfile.json")
	assert.Nil(t, err)

	err = userProfile.SaveProfileAsJson(path)
	assert.Nil(t, err)

	err = userProfile.LoadJsonAsProfile(path)
	assert.Nil(t, err)

	value := userProfile.Agreeableness()
	assert.True(t, value > 1, value)

	value = userProfile.Openness()
	assert.True(t, value > 1, value)

	value = userProfile.Conscientiousness()
	assert.True(t, value > 1, value)

	value = userProfile.Extraversion()
	assert.True(t, value > 1, value)

	value = userProfile.Neuroticism()
	assert.True(t, value > 1, value)

}
