package main

import (
	"fmt"
	"github.com/BachelorGo/core"
	"github.com/BachelorGo/responder"
)

func main() {

	artConsultant := responder.NewArtConsultant()

	bot, err := core.NewSlackBot(artConsultant)
	if err != nil {
		fmt.Println(err)
	}
	bot.Run()
}
