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

func GetAggreeableness(profile personality_insights.Profile) personality_insights.TraitTree {

	value := profile.Tree.Children[0].Children[0].Children[3]

	return value

}

func GetConscientiousness(profile personality_insights.Profile) personality_insights.TraitTree {

	value := profile.Tree.Children[0].Children[0].Children[1]
	return value

}

func GetOpenness(profile personality_insights.Profile) personality_insights.TraitTree {
	value := profile.Tree.Children[0].Children[0].Children[0]
	return value
}
func GetExtraversion(profile personality_insights.Profile) personality_insights.TraitTree {
	value := profile.Tree.Children[0].Children[0].Children[2]
	return value
}
func GetEmotionalStability(profile personality_insights.Profile) personality_insights.TraitTree {
	value := profile.Tree.Children[0].Children[0].Children[4]
	return value
}
