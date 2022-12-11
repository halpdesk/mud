package walk

import (
	"errors"

	"github.com/halpdesk/mud/components/coordinates"
	"github.com/halpdesk/mud/game"
)

var ErrCoordinatesNotExists = errors.New("coordinates does not exist")
var NotPossible = "You cannot go that way"

func New(actor *game.Actor, places []*game.Place) Command {
	return Command{
		actor:  actor,
		places: places,
	}
}

func (c Command) Arity() int {
	return 1
}

type Command struct {
	actor  *game.Actor
	places []*game.Place
}

func (c Command) Execute(args []string) string {
	direction := args[0]
	currentCoordinates := (*(*c.actor).Place()).Coordinates()
	var moveToCoordinates game.Coordinates
	switch direction {
	case "north", "n":
		moveToCoordinates = coordinates.New(currentCoordinates.GetX(), currentCoordinates.GetY()+1)
		break
	case "south", "s":
		moveToCoordinates = coordinates.New(currentCoordinates.GetX(), currentCoordinates.GetY()-1)
		break
	case "east", "e":
		moveToCoordinates = coordinates.New(currentCoordinates.GetX()+1, currentCoordinates.GetY())
		break
	case "west", "w":
		moveToCoordinates = coordinates.New(currentCoordinates.GetX()-1, currentCoordinates.GetY())
		break
	}
	found, room := c.placeCoordinatesExist(moveToCoordinates)
	if !found {
		return NotPossible
	}
	(*c.actor).WalkToPlace(&room)
	return room.CursoryDescription()
}

func (c Command) placeCoordinatesExist(moveToCoordinates game.Coordinates) (bool, game.Place) {
	var exists bool
	var room game.Place
	for _, r := range c.places {
		if moveToCoordinates == (*r).Coordinates() {
			exists = true
			room = *r
		}
	}
	return exists, room
}
