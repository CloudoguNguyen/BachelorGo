package main

import (
	"fmt"
	"github.com/cloudogu/BachelorGo/core"
)

const TOKEN = "fca33d727f5037db12c6039a7efd5d9b"

func main() {
	/*
			re := core.NewRecastClient()
			answer, err := re.GetNextAnswer("tell me a joke", "1")
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(answer[0].Content)

		bot, err := core.NewSlackBot()
		if err != nil {
			fmt.Println(err)
		}
		bot.Run()
				}
	*/
	/*

		file, err := os.Open("resources/contents.json")
		if err != nil {
			fmt.Println(err)
		}
		profile, err := pi.GetProfile(file, "application/json", "en")
		if err != nil {
			fmt.Println(err)

		}
	*/

	pi, err := core.NewPersonalityInsight()
	if err != nil {
		fmt.Println(err)
	}

	err = pi.LoadJsonAsProfile()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(core.GetOpenness(pi.Profile))

}
