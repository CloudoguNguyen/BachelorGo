package service

import (
	"github.com/pkg/errors"
	"github.com/watson-developer-cloud/go-sdk/core"
	"github.com/watson-developer-cloud/go-sdk/personalityinsightsv3"
	"io/ioutil"
)

const (
	watsonUserName = "7bdf7ef4-8c83-4d92-87e6-a03be90b4caf"
	watsonPW       = "OhjpzkGkdkNK"
)

type WatsonPI struct {
	Client *personalityinsightsv3.PersonalityInsightsV3
}

func NewPersonalityInsight() (*WatsonPI, error) {

	client, err := personalityinsightsv3.
		NewPersonalityInsightsV3(&personalityinsightsv3.PersonalityInsightsV3Options{
			Version:  "2017-10-13",
			Username: watsonUserName,
			Password: watsonPW,
		})

	// Check successful instantiation
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create personality insight options ")
	}

	return &WatsonPI{client}, nil

}

func (watson *WatsonPI) GetUserProfile(pathToContent string) (UserProfile, error) {

	userProfile := UserProfile{}

	file, err := ioutil.ReadFile(pathToContent)
	if err != nil {
		return userProfile, errors.Wrapf(err, "failed to open %s", pathToContent)
	}

	profileOptions := watson.Client.
		NewProfileOptions(personalityinsightsv3.ProfileOptions_ContentType_ApplicationJSON)
	profileOptions.SetBody(string(file))
	profileOptions.ContentLanguage = core.StringPtr("en")
	profileOptions.AcceptLanguage = core.StringPtr("en")

	response, err := watson.Client.Profile(profileOptions)
	if err != nil {
		return userProfile, errors.Wrapf(err, "failed to parse profile options")
	}

	profile := watson.Client.GetProfileResult(response)

	userProfile = UserProfile{profile: *profile}

	return userProfile, nil
}
