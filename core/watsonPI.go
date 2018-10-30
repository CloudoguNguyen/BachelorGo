package core

import (
	"encoding/json"
	"fmt"
	"github.com/liviosoares/go-watson-sdk/watson"
	"github.com/liviosoares/go-watson-sdk/watson/personality_insights"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"strconv"
)

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

func (watson *WatsonPI) UpdateProfileWithContent(pathToContent string) error {
	file, err := os.Open(pathToContent)
	if err != nil {
		return errors.Wrapf(err, "failed to open %s", pathToContent)
	}
	profile, err := watson.Client.GetProfile(file, "application/json", "en")
	if err != nil {
		return errors.Wrapf(err, "failed to parse json profile")
	}
	watson.Profile = profile

	return nil
}

func (watson *WatsonPI) GetOpennessValue() int {
	value := watson.Profile.Tree.Children[0].Children[0].Children[0]
	intValue := int(value.Percentage * 100)
	return intValue
}

func (watson *WatsonPI) GetConscientiousnessValue() int {
	value := watson.Profile.Tree.Children[0].Children[0].Children[1]
	intValue := int(value.Percentage * 100)
	return intValue
}
func (watson *WatsonPI) GetExtraversionValue() int {
	value := watson.Profile.Tree.Children[0].Children[0].Children[2]
	intValue := int(value.Percentage * 100)
	return intValue
}

func (watson *WatsonPI) GetAgreeablenessValue() int {
	value := watson.Profile.Tree.Children[0].Children[0].Children[3]
	intValue := int(value.Percentage * 100)
	return intValue
}

func (watson *WatsonPI) GetNeuroticismValue() int {
	value := watson.Profile.Tree.Children[0].Children[0].Children[4]
	intValue := int(value.Percentage * 100)
	return intValue
}

func (watson *WatsonPI) InvalidProfileAsString() string {
	return "Extraversion 999"
}

func (watson *WatsonPI) ProfileAsString() string {
	traits := []string{"Openness", "Conscientiousness", "Extraversion", "Agreeableness", "Neuroticism"}
	var result string

	for index, value := range traits {
		traitValue := int(watson.Profile.Tree.Children[0].Children[0].Children[index].Percentage * 100)
		result += fmt.Sprintf("%s %s; ", value, strconv.Itoa(traitValue))
	}
	return result
}

func (watson *WatsonPI) SaveProfileAsJson(path string) error {
	fo, err := os.Create(path)
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

func (watson *WatsonPI) LoadJsonAsProfile(path string) error {
	jsonFile, err := os.Open(path)
	// if we os.Open returns an error then handle it
	if err != nil {
		return errors.Wrapf(err, "failed to read %s", path)
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return errors.Wrapf(err, "failed to load %s into Json", path)
	}
	json.Unmarshal(byteValue, &watson.Profile)

	return nil

}
