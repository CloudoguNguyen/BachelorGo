package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadAndSaveJsonIntoUserContent(t *testing.T) {
	creator, err := NewMessageCreator("")
	assert.Nil(t, err)

	userContents := UserContents{}

	loadPath := "../resources/1KYCe.json"
	savePath := "../resources/contents2.json"

	err = creator.loadJsonToUserContent(loadPath, &userContents)
	assert.Nil(t, err)

	err = creator.saveUserContentsToJson(savePath, userContents)
	assert.Nil(t, err)

}
