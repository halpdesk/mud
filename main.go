package main

import (
	"fmt"

	"github.com/halpdesk/mud/client"
	"github.com/halpdesk/mud/command"
	"github.com/halpdesk/mud/game"
	"github.com/halpdesk/mud/invoker"
	"github.com/halpdesk/mud/screen"
)

// Server and user sessions
// https://www.youtube.com/watch?v=Eb1Q98PmyLQ

func main() {
	cli := client.New()
	scr := screen.New()
	player := game.NewPlayer("Daniel", 10)
	world := game.NewWorld()

	// Until we qiut
	scr.Clear()
	for {
		// Retrieve the input
		input := cli.Input()

		// Quit
		if input == "quit" {
			fmt.Println("Bye!")
			break
		}

		// Setup command
		cmd := command.New(&world, &player)
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
		scr.WriteLn(result)
	}
}
