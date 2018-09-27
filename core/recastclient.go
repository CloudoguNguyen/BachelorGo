package core

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

	answer := convertMessageToString(response.Messages[0])

	return answer, nil

}

func convertMessageToString(message recast.Component) string {

	stringMessage := fmt.Sprintf("%v", message)

	stringMessage = stringMessage[7 : len(stringMessage)-1]

	return stringMessage

}
