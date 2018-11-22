package service

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestUpdateProfileWIthContent(t *testing.T) {

	pi, err := NewPersonalityInsight()
	assert.Nil(t, err)

	profile, err := pi.GetUserProfile("../resources/test/contents.json")
	assert.Nil(t, err)

	value := profile.Conscientiousness()

	assert.True(t, value > 1, value)
}

func TestSaveAndLoadProfileAsJson(t *testing.T) {

	path := "../resources/test/profile.json"

	pi, err := NewPersonalityInsight()
	assert.Nil(t, err)

	profile, err := pi.GetUserProfile("../resources/test/contents.json")
	assert.Nil(t, err)

	err = profile.SaveProfileAsJson(path)
	assert.Nil(t, err)

	err = profile.LoadJsonAsProfile(path)
	assert.Nil(t, err)

	expected := profile.Agreeableness()

	assert.True(t, expected > 1, expected)

}

func TestProfileAsString(t *testing.T) {

	pi, err := NewPersonalityInsight()
	assert.Nil(t, err)

	profile, err := pi.GetUserProfile("../resources/test/contents.json")
	assert.Nil(t, err)

	result := profile.ProfileAsString()
	fmt.Println(result)

	assert.True(t, strings.Contains(result, "Extraversion 93"))
}
