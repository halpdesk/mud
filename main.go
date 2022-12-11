package main

import (
	"fmt"

	"github.com/halpdesk/mud/components/container"
	"github.com/halpdesk/mud/components/coordinates"
	"github.com/halpdesk/mud/components/item"
	"github.com/halpdesk/mud/components/player"
	"github.com/halpdesk/mud/components/room"
	"github.com/halpdesk/mud/components/world"
	"github.com/halpdesk/mud/core/client"
	"github.com/halpdesk/mud/core/commander"
	"github.com/halpdesk/mud/core/invoker"
	"github.com/halpdesk/mud/core/language"
	"github.com/halpdesk/mud/core/screen"
	"github.com/halpdesk/mud/game"
)

// Server and user sessions
// https://www.youtube.com/watch?v=Eb1Q98PmyLQ

func main() {
	places := setup()
	cli := client.New()
	scr := screen.New()
	player := player.New("Daniel", 10, places[0])
	world := world.New(places)

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
		cmd := commander.New(&world, &player)
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

func setup() []*game.Place {

	var desk = container.New("Desk", "It is wooden and old and has stains on it", language.A, []language.Preposition{language.ON})
	var clock = item.New("Clock", "It's a regular clock, tick tock", language.A)
	var mirror = item.New("Mirror", "There is you", language.A)

	_ = desk.PutObject(&clock, language.ON)

	var start = room.New(
		"Damp cellar",
		"Dark and moist",
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum",
		[]*game.Object{&desk},
		coordinates.New(0, 0),
	)

	var end = room.New(
		"End",
		"This is the end",
		"Enim blandit volutpat maecenas volutpat blandit. A diam sollicitudin tempor id eu nisl nunc mi. Diam sollicitudin tempor id eu nisl. Nec tincidunt praesent semper feugiat nibh sed. Vulputate eu scelerisque felis imperdiet proin fermentum leo",
		[]*game.Object{&mirror},
		coordinates.New(1, 0),
	)

	return []*game.Place{&start, &end}
}
