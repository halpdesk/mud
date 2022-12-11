package main

import (
	"fmt"

	"github.com/halpdesk/mud/components/coordinates"
	"github.com/halpdesk/mud/components/furniture" // ----------------------------------------
	"github.com/halpdesk/mud/components/item"
	"github.com/halpdesk/mud/components/itemcontainer"
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
	intro(scr)
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
		scr.Write80(result)
	}
}

func intro(scr screen.Screen) {
	scr.Write80(screen.Color("Warren", screen.LightGreenFg, screen.BlackBg))
	scr.Write80("------")
	scr.Write80(screen.Color(
		"Warren is a town located in the kingdom of Molag, a kingdom filled with the remnants of ancient magic. It is a dangerous place for many beings. With monsters roaming the countryside and evil sorcerers practicing dark magic, the townsfolk live in fear of the unknown. Corruption has placed a blight on Warren and is slowly spreading. The citizens have banded together to fight back against the corruption.",
		screen.PurpleFg,
		screen.BlackBg,
	))
	scr.Write80("")
	scr.Write80(screen.Color(
		"You wake up on the cellar floor of the Quite Swamp, an inn located in Warren's outer district. What do you do?",
		screen.WhiteFg,
		screen.BlackBg,
	))
}

func setup() []*game.Place {
	var desk = furniture.New("Desk", "An old desk. It is wooden and has ink stains on it", language.A, []language.Preposition{language.ON})
	var boxes = furniture.New("Boxes", "Some wooden boxes in the corner. They are all empty and some are broken", language.SOME, []language.Preposition{language.ON})
	var barrels = furniture.New("Barrels", "Two larger barrels, one ontop of the other. These are typical for storing wine", language.SOME, []language.Preposition{language.ON, language.BEHIND})
	var shelves = furniture.New("Shelves", "A couple of shelves on the far end. There are a couple of boxes and tankards on it, a skillet and some other items, but nothing of interest", language.SOME, []language.Preposition{language.ON})
	var painting = item.New("Painting", "It's a painting depicting the Corruption of Warren", language.A)
	var scale = itemcontainer.New("Scale", "A scale to measure what things weigh", language.A, []language.Preposition{language.ON})
	var sigil = item.New("Stamp", "A stamp with the sigil of The Quiet Swamp Inn", language.A)
	var mirror = item.New("Mirror", "There is you", language.A)
	var torch = item.New("Torch", "It lits up the area", language.A)

	_ = (*desk.Container()).PutObject(&scale, language.ON)
	_ = (*desk.Container()).PutObject(&sigil, language.ON)
	_ = (*barrels.Container()).PutObject(&painting, language.BEHIND)

	var cellar = room.New(
		"Cellar of the Quiet Swamp Inn",
		"A dark and cold place",
		"This cellar holds some supplies for the inn, as well as stock of food and medicines for a long siege. It is partly lit by a torch by the door to the east",
		coordinates.New(0, 0),
	)
	_ = (*cellar.Container()).PutObject(&barrels, language.IN)
	_ = (*cellar.Container()).PutObject(&desk, language.IN)
	_ = (*cellar.Container()).PutObject(&boxes, language.IN)
	_ = (*cellar.Container()).PutObject(&torch, language.IN)
	_ = (*cellar.Container()).PutObject(&shelves, language.IN)

	// var inn = room.New(
	// 	"Inn",
	// 	"Located in a basement in one of Warren's downtown districts",
	// 	"Located in a basement in one of Warren's downtown districts",
	// 	coordinates.New(0, 0),
	// )

	var hallway = room.New(
		"Hallway",
		"A hallway leading to the cellar",
		"On the other end is a staircase leading up to a door",
		coordinates.New(1, 0),
	)
	_ = (*hallway.Container()).PutObject(&mirror, language.IN)

	return []*game.Place{&cellar, &hallway}
}
