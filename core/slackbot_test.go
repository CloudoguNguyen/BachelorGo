package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSlackAppToken(t *testing.T) {
	result, err := getSlackAppToken("../resources/test/testSlackToken")
	assert.Nil(t, err)

	assert.Equal(t, "testToken", result)
}
