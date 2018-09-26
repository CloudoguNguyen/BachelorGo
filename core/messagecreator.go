package core

import (
	"github.com/liviosoares/go-watson-sdk/watson/personality_insights"
	"github.com/pkg/errors"
)

type MessageCreator struct {
	watsonPI *personality_insights.Client
	recast   *RecastClient
}

func NewMessageCreator() (*MessageCreator, error) {

	pi, err := NewPersonalityInsight()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create watson PI")
	}

	recast := NewRecast()

	return &MessageCreator{pi, recast}, nil
}

func (creator *MessageCreator) Response(message string) string {
	return ""
}
