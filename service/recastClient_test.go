package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetIntent(t *testing.T) {
	artConsultantToken := "1fedc8b90ea54efc652b6a42c82de9f2"

	re := NewRecastClient(artConsultantToken)
	intent, err := re.GetIntent("get me art", "testConvID")
	if err != nil {
		assert.Nil(t, err)
	}

	assert.Equal(t, "ask-art", intent.Slug)

}

func TestGetReplay(t *testing.T) {
	artConsultantToken := "1fedc8b90ea54efc652b6a42c82de9f2"

	re := NewRecastClient(artConsultantToken)
	reply, err := re.GetReply("profile not valid", "testConvID")
	if err != nil {
		assert.Nil(t, err)
	}

	assert.NotNil(t, reply)
}
