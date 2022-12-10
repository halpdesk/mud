package say

import (
	"fmt"

	"github.com/halpdesk/mud/game"
)

func New(p game.Player) Command {
	return Command{
		p: p,
	}
}

type Command struct {
	p game.Player
}

func (c Command) Arity() int {
	return 1
}

func (c Command) Execute(args []string) string {
	return fmt.Sprintf("%s says %s", c.p.Name(), args[0])
}
