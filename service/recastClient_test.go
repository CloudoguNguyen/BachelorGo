package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAnswer(t *testing.T) {
	re := NewRecastClient(firstBotToken)
	answer, err := re.GetReplies("Tell me a joke", "1")
	if err != nil {
		assert.Nil(t, err)
	}

	assert.NotNil(t, answer)

}
