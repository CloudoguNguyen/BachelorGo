package main

import (
	"fmt"
	"github.com/cloudogu/BachelorGo/core"
)

const TOKEN = "fca33d727f5037db12c6039a7efd5d9b"

func main() {

	//opts := recast.ConverseOpts{Language: "en"}
	re := core.NewRecastClient()

	rep, err := re.GetReplies("How's the weather?", nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rep)

	/*
		pi, err := core.NewPersonalityInsight()
		if err != nil {
			fmt.Println(err)
		}

		file, err := os.Open("resources/profile.json")
		if err != nil {
			fmt.Println(err)
		}

		profile, err := pi.GetProfile(file, "application/json", "en")
		if err != nil {
			fmt.Println(err)
		}
	*/
	/*bot := core.NewSlackBot()
	bot.Run()*/

}
