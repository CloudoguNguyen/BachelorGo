package core

import "github.com/cloudogu/BachelorGo/service"

type Responder interface {
	GetResponse(message string, conversationID string, watsonPI service.WatsonPI) (string, error)
	GetNewRandomConversationID() string
}
