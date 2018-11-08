package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadAndSaveJsonIntoUserContent(t *testing.T) {
	artConsultant := NewArtConsultant()

	creator, err := NewMessageManager(artConsultant)
	assert.Nil(t, err)

	userContents := UserContents{
		ContentItems: []ContentItem{},
	}

	loadPath := "../resources/test/empty.json"
	savePath := "../resources/test/contents2.json"

	err = creator.loadJsonToUserContent(loadPath, &userContents)
	assert.Nil(t, err)

	contentItem := newContentItem("testMessage")
	userContents.ContentItems = append(userContents.ContentItems, contentItem)

	contentItem = newContentItem("testMessage2")
	userContents.ContentItems = append(userContents.ContentItems, contentItem)

	err = creator.saveUserContentsToJson(savePath, &userContents)
	assert.Nil(t, err)

	err = creator.loadJsonToUserContent(savePath, &userContents)

	assert.Equal(t, "testMessage", userContents.ContentItems[0].Content)

}
