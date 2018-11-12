package core

import (
	"fmt"
	"github.com/cloudogu/BachelorGo/service"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestRecommendArt(t *testing.T) {
	watsonPI := getTestWatsonPI(t)

	consultant := NewArtConsultant()

	s := consultant.recommendArt(*watsonPI)

	assert.True(t, strings.Contains(s, "pop-art"))
	assert.True(t, strings.Contains(s, "kubism"))

}

func getTestWatsonPI(t *testing.T) *service.WatsonPI {
	path := "../resources/test/profile.json"

	watsonPI, err := service.NewPersonalityInsight()
	assert.Nil(t, err)

	watsonPI.LoadJsonAsProfile(path)

	return watsonPI
}

func TestGetIntent(t *testing.T) {
	consultant := NewArtConsultant()

	intent, err := consultant.getIntent("get art", "testConv")

	assert.Nil(t, err)

	fmt.Println(intent)
}

func TestGetResponse(t *testing.T) {
	consultant := NewArtConsultant()
	watsonPI := getTestWatsonPI(t)

	res, err := consultant.GetResponse("get me some art", "testConv", *watsonPI)

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
	watsonPI := getTestWatsonPI(t)

	res, err := consultant.GetResponse("waklalalalalal", "testConv", *watsonPI)

	assert.Nil(t, err)

	assert.Equal(t, "We don't know what you want", res)

}
