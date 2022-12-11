package nocommand

import "fmt"

func New() Command {
	return Command{}
}

func (c Command) Arity() int {
	return 0
}

type Command struct{}

func (c Command) Execute(args []string) string {
	return fmt.Sprintf("Not a valid command")
}
