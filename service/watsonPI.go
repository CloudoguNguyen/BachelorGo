package service

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/watson-developer-cloud/go-sdk/core"
	pi "github.com/watson-developer-cloud/go-sdk/personalityinsightsv3"
	"io/ioutil"
)

const (
	watsonUserName = "7bdf7ef4-8c83-4d92-87e6-a03be90b4caf"
	watsonPW       = "OhjpzkGkdkNK"
)

type WatsonPI struct {
	Client *pi.PersonalityInsightsV3
}

func NewPersonalityInsight() (*WatsonPI, error) {

	client, err := pi.
		NewPersonalityInsightsV3(&pi.PersonalityInsightsV3Options{
			URL:      "https://gateway.watsonplatform.net/personality-insights/api",
			Version:  "2017-10-13",
			Username: watsonUserName,
			Password: watsonPW,
		})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create personality insight options ")
	}

	return &WatsonPI{client}, nil

}

func (watson *WatsonPI) createProfileOption(content pi.Content) *pi.ProfileOptions {
	profileOptions := watson.Client.
		NewProfileOptions(pi.ProfileOptions_ContentType_ApplicationJSON)
	profileOptions.Content = &content
	profileOptions.ContentLanguage = core.StringPtr("en")
	profileOptions.AcceptLanguage = core.StringPtr("en")

	return profileOptions
}

func (watson *WatsonPI) GetUserProfile(pathToContent string) (UserProfile, error) {

	userProfile := UserProfile{}

	file, err := ioutil.ReadFile(pathToContent)
	if err != nil {
		return userProfile, errors.Wrapf(err, "failed to open %s", pathToContent)
	}

	content := new(pi.Content)
	err = json.Unmarshal(file, content)
	if err != nil {
		return userProfile, errors.Wrap(err, "failed to unmarshal json")
	}

	profileOptions := watson.createProfileOption(*content)
	response, err := watson.Client.Profile(profileOptions)
	if err != nil {
		return userProfile, errors.Wrap(err, "failed to parse profile options")
	}

	profile := watson.Client.GetProfileResult(response)
	userProfile = UserProfile{profile: *profile}

	return userProfile, nil
}
