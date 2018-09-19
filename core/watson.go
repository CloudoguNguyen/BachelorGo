package core

import (
	"github.com/liviosoares/go-watson-sdk/watson"
	"github.com/liviosoares/go-watson-sdk/watson/personality_insights"
	"github.com/pkg/errors"
)

func NewPersonalityInsight() (*personality_insights.Client, error) {

	config := watson.Config{
		Credentials: watson.Credentials{
			Username: "7bdf7ef4-8c83-4d92-87e6-a03be90b4caf",
			Password: "OhjpzkGkdkNK",
		},
	}

	client, err := personality_insights.NewClient(config)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create personality client")
	}

	return &client, nil

}
