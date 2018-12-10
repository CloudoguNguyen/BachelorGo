package core

import "github.com/BachelorGo/service"

type Responder interface {
	GetResponse(message string, conversationID string, userProfile *service.UserProfile) (string, error)
}
