package core

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
)

type ContentItem struct {
	Content     string `json:"content"`
	Contenttype string `json:"contenttype"`
	Language    string `json:"language"`
}

type UserContents struct {
	ContentItems []ContentItem `json:"contentItems"`
}

/*
1. Create or read Json file
2. Load it into userContents
3. Add a new contentItem into userContents
4. Delete old Json file
5. Save new userContent into new Json file
*/
func (userContent *UserContents) addMessageToUserContent(message string, jsonPath string) error {

	err := userContent.loadJsonToUserContents(jsonPath)
	if err != nil {
		return errors.Wrapf(err, "failed to load user content %s", jsonPath)
	}

	contentItem := newContentItem(message)
	userContent.ContentItems = append(userContent.ContentItems, contentItem)

	err = userContent.saveUserContentsToJson(jsonPath)
	if err != nil {
		return errors.Wrapf(err, "failed to save user content %s", jsonPath)
	}

	return nil

}

// extract the content of the path into userContent
func (userContent *UserContents) loadJsonToUserContents(path string) error {

	if _, err := os.Stat(path); os.IsNotExist(err) {

		// If file doesn't exist
		jsonFile, err := os.Create(path)
		if err != nil {
			return errors.Wrapf(err, "failed to create %s", path)
		}

		_, err = jsonFile.WriteString("{}")
		if err != nil {
			return errors.Wrapf(err, "failed to write into %s", path)
		}
		defer jsonFile.Close()
	}

	// if we os.Open returns an error then handle it

	jsonFile, err := os.Open(path)
	if err != nil {
		return errors.Wrapf(err, "failed to read %s", path)
	}

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return errors.Wrapf(err, "failed to load %s into Json", path)
	}

	err = json.Unmarshal(byteValue, &userContent)
	if err != nil {
		return errors.Wrapf(err, "failed to unmarshal into Json")
	}
	return nil

}

// save the content of userContents into the path
func (userContent *UserContents) saveUserContentsToJson(path string) error {

	os.Remove(path)

	fo, err := os.Create(path)
	if err != nil {
		return errors.Wrapf(err, "failed to create UserProfile save path")
	}

	defer fo.Close()
	encoder := json.NewEncoder(fo)

	err = encoder.Encode(userContent)
	if err != nil {
		return errors.Wrapf(err, "failed to encode")
	}

	return nil
}
