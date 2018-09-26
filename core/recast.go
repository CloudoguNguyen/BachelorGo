package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

const RequestToken = "2019b5440f2c880dd8ebfc7d2c26df31"

type Recast struct {
	requestToken string
}

func NewRecast() *Recast {

	return &Recast{requestToken: RequestToken}
}

func (recast *Recast) getResponseFromRecastServer(message string, conversationID string) (*http.Response, error) {
	m := Message{
		Content: message,
		Type:    "text",
	}
	data := Payload{
		Message:        m,
		ConversationID: conversationID,
	}

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to marshal %s", data)
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://api.recast.ai/build/v1/dialog", body)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to POST to recast.ai website")
	}
	req.Header.Set("Authorization", "Token "+recast.requestToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to execute the request")
	}

	return resp, nil
}

func (recast *Recast) GetNextAnswer(message string, conversationID string) ([]Message, error) {

	resp, err := recast.getResponseFromRecastServer(message, conversationID)
	if err != nil {
		fmt.Println(err)
	}
	var answerFromRecast RecastResponse

	byteBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(byteBody, &answerFromRecast)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	return answerFromRecast.Results.Messages, nil
}
