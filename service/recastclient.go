package service

import (
	"fmt"
	"github.com/RecastAI/SDK-Golang/recast"
	"github.com/pkg/errors"
)

type RecastClient struct {
	client *recast.RequestClient
}

func NewRecastClient(token string) *RecastClient {

	client := recast.RequestClient{Token: token, Language: "en"}

	return &RecastClient{client: &client}
}

func (rc *RecastClient) GetReplies(message string, conversationID string) (string, error) {

	ops := recast.DialogOpts{Language: "en", ConversationId: conversationID}

	response, err := rc.client.DialogText(message, &ops)
	if err != nil {
		return "", errors.Wrapf(err, "failed to converse text %s", message)
	}

	answer := ""

	for _, message := range response.Messages {

		answer += convertMessageToString(message) + "\n"
	}
	return answer, nil

}

func (rc *RecastClient) GetIntent(message string, conversationID string) (recast.Intent, error) {

	ops := recast.DialogOpts{Language: "en", ConversationId: conversationID}

	response, err := rc.client.DialogText(message, &ops)
	if err != nil {
		return recast.Intent{}, errors.Wrapf(err, "failed to converse text %s", message)
	}

	intent, err := response.Nlp.Intent()
	if err != nil {
		return recast.Intent{}, errors.Wrapf(err, "failed to get intent from text %s", message)
	}

	return intent, err
}

func convertMessageToString(message recast.Component) string {

	stringMessage := fmt.Sprintf("%v", message)
	stringMessage = stringMessage[7 : len(stringMessage)-1]
	return stringMessage
}
