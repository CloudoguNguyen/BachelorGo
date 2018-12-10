package core

import (
	"github.com/BachelorGo/service"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestRecommendArt(t *testing.T) {
	profile := getTestUserProfile(t)

	consultant := NewArtConsultant()

	s := consultant.recommendArt(profile)

	assert.True(t, strings.Contains(s, "pop-art"))
	assert.True(t, strings.Contains(s, "kubism"))

}

func getTestUserProfile(t *testing.T) service.UserProfile {
	path := "../resources/test/profile.json"

	profile := service.UserProfile{}
	profile.LoadJsonAsProfile(path)

	return profile
}

func TestGetIntent(t *testing.T) {
	consultant := NewArtConsultant()

	intent, err := consultant.getIntent("get art", "testConv")

	assert.Nil(t, err)
	assert.Equal(t, "ask-art", intent)
}

func TestGetResponse(t *testing.T) {
	consultant := NewArtConsultant()
	profile := getTestUserProfile(t)

	res, err := consultant.GetResponse("get me some art", "testConv", &profile)

	assert.Nil(t, err)

	assert.True(t, strings.Contains(res, "pop-art"))
	assert.True(t, strings.Contains(res, "kubism"))

}

func TestGetToKnowUser(t *testing.T) {
	consultant := NewArtConsultant()

	res, err := consultant.getToKnowUser("testConv")

	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestGetResponseNoIntentFound(t *testing.T) {
	consultant := NewArtConsultant()
	profile := getTestUserProfile(t)

	res, err := consultant.GetResponse("Wakakakaka", "testConvasdasdasd", &profile)

	assert.Nil(t, err)
	assert.Equal(t, "We have enough information about you now. Please tell us what you want", res)

	consultant.isProfileKnown["testConvasdasdasd"] = true

	res, err = consultant.GetResponse("Wakakakaka", "testConvasdasdasd", &profile)

	assert.Nil(t, err)
	assert.Equal(t, "We don't know what you want", res)

}
