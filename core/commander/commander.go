package commander

import (
	"errors"
	"strings"

	"github.com/halpdesk/mud/commands/look"
	"github.com/halpdesk/mud/commands/say"
	"github.com/halpdesk/mud/commands/take"
	"github.com/halpdesk/mud/commands/walk"
	"github.com/halpdesk/mud/core/contracts"
	"github.com/halpdesk/mud/game"
)

var ErrNotValidCommand = errors.New("not a valid command")

func New(world *game.World, actor *game.Actor) Command {
	return Command{
		world: world,
		actor: actor,
	}
}

type Command struct {
	world *game.World
	actor *game.Actor
}

func (c Command) GetCommandAndArgs(input string) (contracts.Command, []string, error) {
	var cmd contracts.Command
	var err error
	base, args := c.interpret(input)
	// fmt.Printf("-- Command was: %s with %+v", base, args)
	switch base {
	case "go":
		cmd = walk.New(c.actor, (*c.world).Places())
		break
	case "say":
		cmd = say.New(c.actor)
		break
	case "look":
		cmd = look.New(c.actor, (*c.actor).Place())
		break
	case "take":
		cmd = take.New(c.actor, (*c.actor).Place())
		break
	default:
		err = ErrNotValidCommand
		break
	}
	return cmd, args, err
}

func (c Command) interpret(input string) (string, []string) {
	pieces := strings.Split(input, " ")
	var args []string
	base := pieces[0]
	if len(pieces) > 0 {
		args = pieces[1:]
	}
	return c.aliaser(base, args)
}

func (c Command) aliaser(base string, args []string) (string, []string) {
	if base == "n" || base == "north" {
		return "go", []string{"north"}
	}
	if base == "s" || base == "south" {
		return "go", []string{"south"}
	}
	if base == "e" || base == "east" {
		return "go", []string{"east"}
	}
	if base == "w" || base == "west" {
		return "go", []string{"west"}
	}
	if base == "l" {
		return "look", args
	}
	return base, args
}
