package core

import "github.com/cloudogu/BachelorGo/service"

type Responder interface {
	GetResponse(message string, conversationID string, userProfile *service.UserProfile) (string, error)
	GetNewRandomConversationID() string
}
