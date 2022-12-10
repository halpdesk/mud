package contracts

// type CommandFunc func(arg string) string

type Command interface {
	Execute(args []string) string
	Arity() int
}
