package core

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type MessageManager struct {
	watsonPI    *WatsonPI
	recast      *RecastClient
	enoughWords bool
}

func NewMessageCreator(recastToken string) (*MessageManager, error) {

	watsonPI, err := NewPersonalityInsight()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create watson PI")
	}
	recastClient := NewRecastClient(recastToken)

	return &MessageManager{watsonPI, recastClient, true}, nil
}

func (manager *MessageManager) Response(message string, conversationID string) (string, error) {

	path := "resources/" + conversationID + ".json"

	err := manager.addMessageIntoConversationJson(message, path)
	if err != nil {
		return "", errors.Wrapf(err, "failed to add message to json with %s", conversationID)
	}

	err = manager.watsonPI.UpdateProfileWithContent(path)
	if err != nil {
		if strings.Contains(err.Error(), "less than the minimum number of words required") {
			manager.enoughWords = false
			return "We need atleast 100 words from you to analyse your personality, please tell us more about you", nil
		}
		return "", errors.Wrapf(err, "failed update profile in conversation %s", conversationID)
	}

	if manager.enoughWords == false {
		manager.enoughWords = true
		return "We have enough words from you now, please tell us what you want", nil
	}

	fmt.Println(manager.watsonPI.ProfileAsString())
	messageForRecast := message + " extraversion " + strconv.Itoa(manager.watsonPI.GetExtraversionValue())
	fmt.Println("Message to recast: " + messageForRecast)

	answer, err := manager.recast.GetReplies(messageForRecast, conversationID)
	if err != nil {
		return "", errors.Wrapf(err, "failed to get reply with the messsage %s", message)
	}
	return answer, nil
}

func (manager *MessageManager) NewConversationID() string {

	newID := manager.recast.getNewRandomConversationID()

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
		return errors.Wrapf(err, "failed to create Profile save path")
	}

	defer fo.Close()
	encoder := json.NewEncoder(fo)

	err = encoder.Encode(userContent)
	if err != nil {
		return errors.Wrapf(err, "failed to create encode")
	}

	return nil
}
