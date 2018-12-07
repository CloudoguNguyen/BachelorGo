package core

import (
	"encoding/json"
	"github.com/BachelorGo/service"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"strings"
)

type MessageManager struct {
	watsonPI    *service.WatsonPI
	responder   Responder
	enoughWords bool
}

const profile_not_valid = "profile not valid"

func NewMessageManager(responder Responder) (*MessageManager, error) {

	watsonPI, err := service.NewPersonalityInsight()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create watson PI")
	}

	return &MessageManager{watsonPI, responder, true}, nil
}

func (manager *MessageManager) Response(message string, conversationID string) (string, error) {

	path := "resources/conversations/" + conversationID + ".json"

	err := manager.addMessageIntoConversationJson(message, path)
	if err != nil {
		return "", errors.Wrapf(err, "failed to add message to json with %s", conversationID)
	}

	profile, err := manager.getUserProfile(path)
	if err != nil {
		if strings.Contains(err.Error(), "less than the minimum number of words required") {
			message = profile_not_valid
		} else {
			return "", errors.Wrapf(err, "failed to update profile in conversation %s", conversationID)
		}
	}

	answer, err := manager.responder.GetResponse(message, conversationID, &profile)
	if err != nil {
		return "", errors.Wrapf(err, "failed to get reply with the message %s", message)
	}
	return answer, nil
}

func (manager *MessageManager) getUserProfile(path string) (service.UserProfile, error) {

	profile, err := manager.watsonPI.GetUserProfile(path)
	if err != nil {
		return service.UserProfile{}, errors.Wrapf(err, "failed update profile in conversation")
	}
	return profile, nil
}

func (manager *MessageManager) NewConversationID() string {

	newID := manager.responder.GetNewRandomConversationID()

	return newID
}

/*
1. Create/read Json file
2. Load it into userContent
3. Add contentItem into userContent
4. Delete old Json file
5. Save new userContent into new JsonFile
*/
func (manager *MessageManager) addMessageIntoConversationJson(message string, jsonPath string) error {

	userContent := UserContents{}
	err := manager.loadJsonToUserContent(jsonPath, &userContent)
	if err != nil {
		return errors.Wrapf(err, "failed to load user content %s", jsonPath)
	}

	contentItem := newContentItem(message)
	userContent.ContentItems = append(userContent.ContentItems, contentItem)

	err = manager.saveUserContentsToJson(jsonPath, &userContent)
	if err != nil {
		return errors.Wrapf(err, "failed to save user content %s", jsonPath)
	}

	return nil

}

func (manager *MessageManager) loadJsonToUserContent(path string, content *UserContents) error {

	if _, err := os.Stat(path); os.IsNotExist(err) {

		// If file doenst exist
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

	err = json.Unmarshal(byteValue, &content)
	if err != nil {
		return errors.Wrapf(err, "failed to unmarshal into Json")
	}
	return nil

}

func (manager *MessageManager) saveUserContentsToJson(path string, userContent *UserContents) error {

	os.Remove(path)

	fo, err := os.Create(path)
	if err != nil {
		return errors.Wrapf(err, "failed to create UserProfile save path")
	}

	defer fo.Close()
	encoder := json.NewEncoder(fo)

	err = encoder.Encode(userContent)
	if err != nil {
		return errors.Wrapf(err, "failed to create encode")
	}

	return nil
}
