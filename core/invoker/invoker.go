package invoker

import (
	"errors"

	"github.com/halpdesk/mud/core/contracts"
)

var ErrInvalidArity = errors.New("not enough arguments")

type Invoker struct {
	c contracts.Command
}

func New(c contracts.Command) Invoker {
	return Invoker{
		c: c,
	}
}

func (i Invoker) Do(args []string) (string, error) {
	// fmt.Printf("Command is: %T", i.c)
	if len(args) < i.c.Arity() {
		return "", ErrInvalidArity
	}
	return i.c.Execute(args), nil
}
