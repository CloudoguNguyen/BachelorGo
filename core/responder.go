package core

import "github.com/liviosoares/go-watson-sdk/watson/personality_insights"

type Responder interface {
	GetResponse(message string, conversationID string, profile personality_insights.Profile) (string, error)
	GetNewRandomConversationID() string
}
