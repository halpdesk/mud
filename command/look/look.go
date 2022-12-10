package look

import (
	"errors"

	"github.com/halpdesk/randomwalker/game"
)

var ErrCoordinatesNotExists = errors.New("coordinates does not exist")
var NotPossible = "You cannot go that way"

func New(room game.Room) Command {
	return Command{
		room: room,
	}
}

func (c Command) Arity() int {
	return 0
}

type Command struct {
	room game.Room
}

func (c Command) Execute(args []string) string {
	return c.room.Description()
}
