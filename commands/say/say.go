package say

import (
	"fmt"

	"github.com/halpdesk/mud/game"
)

func New(p *game.Actor) Command {
	return Command{
		p: p,
	}
}

type Command struct {
	p *game.Actor
}

func (c Command) Arity() int {
	return 1
}

func (c Command) Execute(args []string) string {
	return fmt.Sprintf("%s said %s", (*c.p).Name(), args[0])
}
