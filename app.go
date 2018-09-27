package main

import (
	"fmt"
	"github.com/cloudogu/BachelorGo/core"
)

func main() {

	bot, err := core.NewSlackBot()
	if err != nil {
		fmt.Println(err)
	}
	bot.Run()
}
