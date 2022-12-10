package walk

import (
	"errors"
	"fmt"

	"github.com/halpdesk/randomwalker/game"
)

var ErrCoordinatesNotExists = errors.New("coordinates does not exist")
var NotPossible = "You cannot go that way"

func New(player *game.Player, rooms []*game.Room) Command {
	return Command{
		player: player,
		rooms:  rooms,
	}
}

func (c Command) Arity() int {
	return 1
}

type Command struct {
	player *game.Player
	rooms  []*game.Room
}

func (c Command) Execute(args []string) string {
	direction := args[0]
	moveToCoordinates := c.player.Room().Coordinates()
	switch direction {
	case "north":
		moveToCoordinates.Y = moveToCoordinates.Y + 1
		break
	case "south":
		moveToCoordinates.Y = moveToCoordinates.Y - 1
		break
	case "east":
		moveToCoordinates.X = moveToCoordinates.X + 1
		break
	case "west":
		moveToCoordinates.X = moveToCoordinates.X - 1
		break
	}
	found, room := c.roomCoordinatesExist(moveToCoordinates)
	if !found {
		return NotPossible
	}
	c.player.WalkToRoom(&room)
	return fmt.Sprintf("You went %s", args[0])
}

func (c Command) roomCoordinatesExist(moveToCoordinates game.Coordinates) (bool, game.Room) {
	var exists bool
	var room game.Room
	for _, r := range c.rooms {
		if moveToCoordinates == r.Coordinates() {
			exists = true
			room = *r
		}
	}
	return exists, room
}
