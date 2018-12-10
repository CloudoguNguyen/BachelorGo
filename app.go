package main

import (
	"fmt"
	"github.com/BachelorGo/core"
)

func main() {

	artConsultant := core.NewArtConsultant()

	bot, err := core.NewSlackBot(artConsultant)
	if err != nil {
		fmt.Println(err)
	}
	bot.Run()
}
