package core

import (
	"fmt"
	"github.com/RecastAI/SDK-Golang/recast"
	"github.com/pkg/errors"
	"math/rand"
	"time"
)

type RecastClient struct {
	client *recast.RequestClient
}

const firstBotToken = "2019b5440f2c880dd8ebfc7d2c26df31"
const secondBotToken = "e16b673cc84ab7b5d490115dedfe7d71"

func NewRecastClient(token string) *RecastClient {

	if token == "" {
		token = firstBotToken
	}

	client := recast.RequestClient{Token: token, Language: "en"}

	return &RecastClient{client: &client}
}

func (rc *RecastClient) GetReplies(message string, conversationID string) (string, error) {

	ops := recast.DialogOpts{Language: "en", ConversationId: conversationID}

	response, err := rc.client.DialogText(message, &ops)
	if err != nil {
		return "", errors.Wrapf(err, "failed to converse text %s", message)
	}

	if len(response.Messages) > 0 {
		answer := convertMessageToString(response.Messages[0])
		return answer, nil
	}

	return "I dont understand it yet", nil

}

func convertMessageToString(message recast.Component) string {

	stringMessage := fmt.Sprintf("%v", message)
	stringMessage = stringMessage[7 : len(stringMessage)-1]

	return stringMessage

}

func (rc *RecastClient) getNewConversationID() string {

	return newRandomConversationID()
}

func newRandomConversationID() string {
	rand.Seed(time.Now().UnixNano())

	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	b := make([]rune, 10)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
