package core

import (
	"fmt"
	"github.com/nlopes/slack"
	"strings"
)

type SlackBot struct {
	token string
}

func NewSlackBot() *SlackBot {
	return &SlackBot{token: "xoxb-438453325860-438070557617-CviJFdimezMGe8FM04MwfO5a"}
}

func (bot *SlackBot) Run() {
	api := slack.New(bot.token)
	rtm := api.NewRTM()
	go rtm.ManageConnection()
Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			fmt.Print("Event Received: ")
			switch ev := msg.Data.(type) {
			case *slack.ConnectedEvent:
				fmt.Println("Connection counter:", ev.ConnectionCount)

			case *slack.MessageEvent:
				fmt.Printf("Message: %v\n", ev.Text)
				info := rtm.GetInfo()
				prefix := fmt.Sprintf("<@%s> ", info.User.ID)

				bot.Respond(rtm, ev, prefix)

			case *slack.RTMError:
				fmt.Printf("Error: %s\n", ev.Error())

			case *slack.InvalidAuthEvent:
				fmt.Printf("Invalid credentials")
				break Loop

			default:
				//Take no action
			}
		}
	}

}

func (bot *SlackBot) Respond(rtm *slack.RTM, msg *slack.MessageEvent, prefix string) {
	var response string
	text := msg.Text
	text = strings.TrimPrefix(text, prefix)
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)

	response = "Im working"
	rtm.SendMessage(rtm.NewOutgoingMessage(response, msg.Channel))

}
