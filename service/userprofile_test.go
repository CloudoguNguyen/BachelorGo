package service

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestProfileAsString(t *testing.T) {

	userProfile := UserProfile{}

	err := userProfile.LoadJsonAsProfile("../resources/test/profile.json")
	assert.Nil(t, err)

	result := userProfile.ProfileAsString()
	fmt.Println(result)

	assert.True(t, strings.Contains(result, "Extraversion 93"))
}

func TestSaveAndLoadProfileAsJson(t *testing.T) {

	path := "../resources/test/profile2.json"

	userProfile := UserProfile{}

	err := userProfile.LoadJsonAsProfile("../resources/test/profile.json")
	assert.Nil(t, err)

	err = userProfile.SaveProfileAsJson(path)
	assert.Nil(t, err)

	err = userProfile.LoadJsonAsProfile(path)
	assert.Nil(t, err)

	value := userProfile.Agreeableness()

	assert.True(t, value > 1, value)

}
