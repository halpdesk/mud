package main

import (
	"fmt"

	// ----------------------------------------
	"github.com/halpdesk/mud/components/player"
	"github.com/halpdesk/mud/components/world"
	"github.com/halpdesk/mud/core/client"
	"github.com/halpdesk/mud/core/commander"
	"github.com/halpdesk/mud/core/invoker"
	"github.com/halpdesk/mud/core/screen"
	"github.com/halpdesk/mud/data"
)

// Server and user sessions
// https://www.youtube.com/watch?v=Eb1Q98PmyLQ

func main() {
	cli := client.New()
	scr := screen.New()

	places := data.Setup()
	player := player.New("Daniel", 10, places[0])
	world := world.New(places)

	cmd := commander.New(&world, &player)

	// Until we qiut
	scr.Clear()
	data.Intro(scr)
	for {
		// Retrieve the input
		input := cli.Input()

		// Quit command
		if input == "quit" {
			fmt.Println("Bye!")
			break
		}

		// Parse command
		invokeCmd, args, err := cmd.GetCommandAndArgs(input)
		if err != nil {
			scr.WriteLn("%s %s", screen.Color("ERROR:", screen.WhiteFg, screen.RedBg), err.Error())
			continue
		}

		// Invoke
		result, err := invoker.New(invokeCmd).Do(args)
		if err != nil {
			scr.WriteLn("%s %s", screen.Color("ERROR:", screen.WhiteFg, screen.RedBg), err.Error())
			continue
		}
		scr.Write80(result)
	}
}
