package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadAndSaveJsonIntoUserContent(t *testing.T) {

	userContents := UserContents{
		ContentItems: []ContentItem{},
	}

	loadPath := "../resources/test/empty.json"
	savePath := "../resources/test/contents2.json"

	err := userContents.loadJsonToUserContents(loadPath)
	assert.Nil(t, err)

	contentItem := newContentItem("testMessage")
	userContents.ContentItems = append(userContents.ContentItems, contentItem)

	contentItem = newContentItem("testMessage2")
	userContents.ContentItems = append(userContents.ContentItems, contentItem)

	err = userContents.saveUserContentsToJson(savePath)
	assert.Nil(t, err)

	err = userContents.loadJsonToUserContents(savePath)

	assert.Equal(t, "testMessage2", userContents.ContentItems[1].Content)

}
