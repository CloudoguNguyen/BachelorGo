package core

import (
	"encoding/json"
	"github.com/liviosoares/go-watson-sdk/watson"
	"github.com/liviosoares/go-watson-sdk/watson/personality_insights"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
)

func NewPersonalityInsight() (*personality_insights.Client, error) {

	//ToDo username and pw into own file?
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

func SaveJsonFile(v interface{}, path string) error {
	fo, err := os.Create(path)
	if err != nil {
		return errors.Wrapf(err, "failed to create %s", path)
	}

	defer fo.Close()
	encoder := json.NewEncoder(fo)

	err = encoder.Encode(v)
	if err != nil {
		return errors.Wrapf(err, "failed to create encode")
	}
	return nil
}

func LoadJsonProfile(path string) (*personality_insights.Profile, error) {
	jsonFile, err := os.Open(path)
	// if we os.Open returns an error then handle it
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read %s", path)
	}
	defer jsonFile.Close()

	var profile personality_insights.Profile

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to load %s into Json", path)
	}
	json.Unmarshal(byteValue, &profile)

	return &profile, nil
}
