package service

import (
	"github.com/liviosoares/go-watson-sdk/watson"
	"github.com/liviosoares/go-watson-sdk/watson/personality_insights"
	"github.com/pkg/errors"
	"os"
)

const (
	watsonUserName = "7bdf7ef4-8c83-4d92-87e6-a03be90b4caf"
	watsonPW       = "OhjpzkGkdkNK"
)

type WatsonPI struct {
	Client personality_insights.Client
}

func NewPersonalityInsight() (*WatsonPI, error) {

	config := watson.Config{
		Credentials: watson.Credentials{
			Username: watsonUserName,
			Password: watsonPW,
		},
	}

	client, err := personality_insights.NewClient(config)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create personality client")
	}

	return &WatsonPI{client}, nil

}

func (watson *WatsonPI) GetUserProfile(pathToContent string) (UserProfile, error) {

	userProfile := UserProfile{}

	file, err := os.Open(pathToContent)
	if err != nil {
		return userProfile, errors.Wrapf(err, "failed to open %s", pathToContent)
	}

	profile, err := watson.Client.GetProfile(file, "application/json", "en")
	if err != nil {
		return userProfile, errors.Wrapf(err, "failed to get profile from json file")
	}

	userProfile = UserProfile{profile: profile}

	return userProfile, nil
}
