package core

import (
	"fmt"
	"github.com/pkg/errors"
)

type MessageCreator struct {
	watsonPI *WatsonPI
	recast   *RecastClient
}

func NewMessageCreator() (*MessageCreator, error) {

	pi, err := NewPersonalityInsight()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create watson PI")
	}
	recastClient := NewRecastClient()

	return &MessageCreator{pi, recastClient}, nil
}

func (creator *MessageCreator) Response(message string, conversationID string) string {

	answer, err := creator.recast.GetNextAnswer(message, conversationID)
	if err != nil {
		fmt.Println(err)
	}
	return answer[0].Content
}

//ToDo add message into contents2.json
