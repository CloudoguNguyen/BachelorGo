package core

import (
	"github.com/cloudogu/BachelorGo/service"
	"github.com/liviosoares/go-watson-sdk/watson/personality_insights"
)

type ArtConsultant struct {
	recast *service.RecastClient
}

func NewArtConsultant(recastToken string) *ArtConsultant {

	recast := service.NewRecastClient(recastToken)

	return &ArtConsultant{recast}
}

func (ac *ArtConsultant) GetResponse(message string, conversationID string, profile personality_insights.Profile) (string, error) {

	return "", nil
}

func (ac *ArtConsultant) GetNewRandomConversationID() string {

	return ""
}
