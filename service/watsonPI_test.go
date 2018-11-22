package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdateProfileWIthContent(t *testing.T) {

	pi, err := NewPersonalityInsight()
	assert.Nil(t, err)

	profile, err := pi.GetUserProfile("../resources/test/contents.json")
	assert.Nil(t, err)

	value := profile.Conscientiousness()

	assert.True(t, value > 1, value)
}
