package core

import (
	"fmt"
	"github.com/RecastAI/SDK-Golang/recast"
	"github.com/pkg/errors"
)

const RequestToken = "e16b673cc84ab7b5d490115dedfe7d71"

type RecastClient struct {
	client *recast.RequestClient
}

func NewRecastClient() *RecastClient {

	client := recast.RequestClient{Token: RequestToken, Language: "en"}

	return &RecastClient{client: &client}
}

func (rc *RecastClient) GetReplies(message string, opts *recast.ConverseOpts) ([]string, error) {

	response, err := rc.client.ConverseText(message, opts)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to converse text %s", message)
	}

	fmt.Println(response)

	fmt.Println(response.Action)
	fmt.Println(response.Language)

	return response.Replies, nil

}
