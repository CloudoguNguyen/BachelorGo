package core

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
)

type MessageCreator struct {
	watsonPI *WatsonPI
	recast   *RecastClient
}

func NewMessageCreator(token string) (*MessageCreator, error) {

	pi, err := NewPersonalityInsight()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create watson PI")
	}
	recastClient := NewRecastClient(token)

	return &MessageCreator{pi, recastClient}, nil
}

//ToDo make watson read it

func (creator *MessageCreator) Response(message string, conversationID string) (string, error) {

	err := creator.addMessageIntoJson(message, conversationID)
	if err != nil {
		return "", errors.Wrapf(err, "failed to add message to json with %s", conversationID)
	}

	answer, err := creator.recast.GetReplies(message, conversationID)
	if err != nil {
		return "", errors.Wrapf(err, "failed to get reply with the messsage %s", message)
	}
	return answer, nil
}

func (creator *MessageCreator) NewConversationID() string {

	newID := creator.recast.getNewConversationID()

	return newID
}

/*
1. Create/read Json file
2. Load it into userContent
3. Add contentItem into userContent
4. Delete old Json file
5. Save new userContent into new JsonFile
*/
func (creator *MessageCreator) addMessageIntoJson(message string, conversationID string) error {

	path := "resources/" + conversationID + ".json"

	userContent := UserContents{}
	err := creator.loadJsonToUserContent(path, &userContent)
	if err != nil {
		return errors.Wrapf(err, "failed to load user content %s", path)
	}

	contentItem := newContentItem(message)
	userContent.ContentItems = append(userContent.ContentItems, contentItem)

	err = creator.saveUserContentsToJson(path, userContent)
	if err != nil {
		return errors.Wrapf(err, "failed to save user content %s", path)
	}

	return nil

}

func (creator *MessageCreator) loadJsonToUserContent(path string, content *UserContents) error {

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

func (creator *MessageCreator) saveUserContentsToJson(path string, userContent UserContents) error {

	os.Remove(path)

	fo, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return errors.Wrapf(err, "failed to create path %s", path)
	}

	defer fo.Close()
	encoder := json.NewEncoder(fo)

	err = encoder.Encode(userContent)
	if err != nil {
		return errors.Wrapf(err, "failed to encode")
	}
	return nil
}
