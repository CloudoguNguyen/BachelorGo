package core

import (
	"fmt"
	"github.com/nlopes/slack"
	"github.com/pkg/errors"
)

type SlackBot struct {
	token   string
	client  *slack.Client
	rtm     *slack.RTM
	creator *MessageCreator
}

func NewSlackBot() (*SlackBot, error) {

	token := "xoxb-438453325860-438070557617-CviJFdimezMGe8FM04MwfO5a"
	client := slack.New(token)
	rtm := client.NewRTM()
	creator, err := NewMessageCreator()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create MessageCreator")
	}

	return &SlackBot{token: token, client: client, rtm: rtm, creator: creator}, nil
}

func (bot *SlackBot) Run() {

	go bot.rtm.ManageConnection()
	for {
		select {
		case message := <-bot.rtm.IncomingEvents:
			fmt.Print("Event Received: ")

			switch event := message.Data.(type) {
			case *slack.ConnectedEvent:
				fmt.Println("Connection counter:", event.ConnectionCount)

			case *slack.MessageEvent:
				fmt.Printf("Message: %v\n", event.Text)

				bot.Respond(event)

			case *slack.RTMError:
				fmt.Printf("Error: %s\n", event.Error())

			case *slack.InvalidAuthEvent:
				fmt.Printf("Invalid credentials")
				break
			}
		}
	}

}

func (bot *SlackBot) Respond(msg *slack.MessageEvent) {
	var response string
	text := msg.Text

	response = bot.creator.Response(text)

	bot.rtm.SendMessage(bot.rtm.NewOutgoingMessage(response, msg.Channel))

}
