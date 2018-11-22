package service

import (
	"encoding/json"
	"fmt"
	"github.com/liviosoares/go-watson-sdk/watson/personality_insights"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"strconv"
)

type UserProfile struct {
	profile personality_insights.Profile
}

func (profile *UserProfile) Openness() int {
	value := profile.profile.Tree.Children[0].Children[0].Children[0]
	intValue := int(value.Percentage * 100)
	return intValue
}

func (profile *UserProfile) Conscientiousness() int {
	value := profile.profile.Tree.Children[0].Children[0].Children[1]
	intValue := int(value.Percentage * 100)
	return intValue
}
func (profile *UserProfile) Extraversion() int {
	value := profile.profile.Tree.Children[0].Children[0].Children[2]
	intValue := int(value.Percentage * 100)
	return intValue
}

func (profile *UserProfile) Agreeableness() int {
	value := profile.profile.Tree.Children[0].Children[0].Children[3]
	intValue := int(value.Percentage * 100)
	return intValue
}

func (profile *UserProfile) Neuroticism() int {
	value := profile.profile.Tree.Children[0].Children[0].Children[4]
	intValue := int(value.Percentage * 100)
	return intValue
}

func (profile *UserProfile) SaveProfileAsJson(path string) error {
	fo, err := os.Create(path)
	if err != nil {
		return errors.Wrapf(err, "failed to create UserProfile save path")
	}

	defer fo.Close()
	encoder := json.NewEncoder(fo)

	err = encoder.Encode(profile.profile)
	if err != nil {
		return errors.Wrapf(err, "failed to create encode")
	}
	return nil
}

func (profile *UserProfile) LoadJsonAsProfile(path string) error {
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
	json.Unmarshal(byteValue, &profile.profile)

	return nil
}

func (profile *UserProfile) ProfileAsString() string {
	traits := []string{"Openness", "Conscientiousness", "Extraversion", "Agreeableness", "Neuroticism"}
	var result string

	for index, value := range traits {
		traitValue := int(profile.profile.Tree.Children[0].Children[0].Children[index].Percentage * 100)
		result += fmt.Sprintf("%s %s; ", value, strconv.Itoa(traitValue))
	}
	return result
}

func (profile *UserProfile) InvalidProfileAsString() string {
	return "Extraversion 999"
}
