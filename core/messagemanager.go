package core

import (
	"github.com/BachelorGo/responder"
	"github.com/BachelorGo/service"
	"github.com/pkg/errors"
	"math/rand"
	"strings"
	"time"
)

type MessageManager struct {
	watsonPI    *service.WatsonPI
	responder   responder.Responder
	enoughWords bool
}

func NewMessageManager(responder responder.Responder) (*MessageManager, error) {

	watsonPI, err := service.NewPersonalityInsight()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create watson PI")
	}

	return &MessageManager{watsonPI, responder, true}, nil
}

func (manager *MessageManager) Response(message string, conversationID string) (string, error) {

	path := "resources/conversations/" + conversationID + ".json"

	userContent := UserContents{}

	err := userContent.addMessageToUserContent(message, path)
	if err != nil {
		return "", errors.Wrapf(err, "failed to add message to json with %s", conversationID)
	}

	profile, err := manager.getUserProfile(path)
	if err != nil {
		if strings.Contains(err.Error(), "less than the minimum number of words required") {
			message = responder.ProfileNotValid
		} else {
			return "", errors.Wrapf(err, "failed to update profile in conversation %s", conversationID)
		}
	}

	answer, err := manager.responder.GetResponse(message, conversationID, &profile)
	if err != nil {
		return "", errors.Wrapf(err, "failed to get reply with the message")
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

func (manager *MessageManager) NewRandomConversationID() string {

	rand.Seed(time.Now().UnixNano())

	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	b := make([]rune, 10)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func newContentItem(message string) ContentItem {
	return ContentItem{
		Content:     message,
		Contenttype: "text/plain",
		Language:    "en",
	}
}
