package core

import (
	"fmt"
	"github.com/RecastAI/SDK-Golang/recast"
	"github.com/pkg/errors"
)

const firstBotToken = "2019b5440f2c880dd8ebfc7d2c26df31"
const secondBotToken = "e16b673cc84ab7b5d490115dedfe7d71"

type RecastClient struct {
	client *recast.RequestClient
}

func NewRecastClient() *RecastClient {

	client := recast.RequestClient{Token: firstBotToken, Language: "en"}

	return &RecastClient{client: &client}
}

func (rc *RecastClient) GetReplies(message string, conversationID string) (string, error) {

	ops := recast.DialogOpts{Language: "en", ConversationId: conversationID}

	response, err := rc.client.DialogText(message, &ops)
	if err != nil {
		return "", errors.Wrapf(err, "failed to converse text %s", message)
	}

	answer := convertMessageToString(response.Messages[0])

	return answer, nil

}

func convertMessageToString(message recast.Component) string {

	stringMessage := fmt.Sprintf("%v", message)

	stringMessage = stringMessage[7 : len(stringMessage)-1]

	return stringMessage

}
