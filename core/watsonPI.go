package core

import (
	"encoding/json"
	"fmt"
	"github.com/liviosoares/go-watson-sdk/watson"
	"github.com/liviosoares/go-watson-sdk/watson/personality_insights"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
)

const profileSavePath = "/home/tnguyen/GolandProjects/src/github.com/cloudogu/BachelorGo/resources/profile.json"

type WatsonPI struct {
	Client  personality_insights.Client
	Profile personality_insights.Profile
}

func NewPersonalityInsight() (*WatsonPI, error) {

	//ToDo username and pw into own file?
	config := watson.Config{
		Credentials: watson.Credentials{
			Username: "7bdf7ef4-8c83-4d92-87e6-a03be90b4caf",
			Password: "OhjpzkGkdkNK",
		},
	}

	client, err := personality_insights.NewClient(config)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create personality Client")
	}

	var profile personality_insights.Profile

	return &WatsonPI{client, profile}, nil

}

func (watson *WatsonPI) updateProfileWithContent(pathToContent string) {
	file, err := os.Open(pathToContent)
	if err != nil {
		fmt.Println(err)
	}
	profile, err := watson.Client.GetProfile(file, "application/json", "en")
	if err != nil {
		fmt.Println(err)
	}

	watson.Profile = profile
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

func (watson *WatsonPI) SaveProfileAsJson() error {
	fo, err := os.Create(profileSavePath)
	if err != nil {
		return errors.Wrapf(err, "failed to create Profile save path")
	}

	defer fo.Close()
	encoder := json.NewEncoder(fo)

	err = encoder.Encode(watson.Profile)
	if err != nil {
		return errors.Wrapf(err, "failed to create encode")
	}
	return nil
}

func (watson *WatsonPI) LoadJsonAsProfile() error {
	jsonFile, err := os.Open(profileSavePath)
	// if we os.Open returns an error then handle it
	if err != nil {
		return errors.Wrapf(err, "failed to read %s", profileSavePath)
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return errors.Wrapf(err, "failed to load %s into Json", profileSavePath)
	}
	json.Unmarshal(byteValue, &watson.Profile)

	return nil

}
