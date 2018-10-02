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

	value := pi.GetConscientiousnessValue()

	assert.True(t, value > 1, value)
}

func TestSaveAndLoadProfileAsJson(t *testing.T) {

	path := "../resources/test/profile.json"

	pi, err := NewPersonalityInsight()
	assert.Nil(t, err)

	err = pi.updateProfileWithContent("../resources/test/contents.json")
	assert.Nil(t, err)

	err = pi.SaveProfileAsJson(path)
	assert.Nil(t, err)

	err = pi.LoadJsonAsProfile(path)
	assert.Nil(t, err)

	expected := pi.GetAgreeablenessValue()

	assert.True(t, expected > 1, expected)

}
