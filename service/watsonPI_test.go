package service

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdateProfileWIthContent(t *testing.T) {

	pi, err := NewPersonalityInsight()
	assert.Nil(t, err)

	profile, err := pi.GetUserProfile("../resources/test/contents.json")
	assert.Nil(t, err)

	fmt.Println(profile)

}
