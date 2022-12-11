package look

import (
	"github.com/halpdesk/mud/core/language"
	"github.com/halpdesk/mud/game"
)

var NothingFound = "There is no such thing here"

func New(actor *game.Actor, place *game.Place) Command {
	return Command{
		place: place,
		actor: actor,
	}
}

func (c Command) Arity() int {
	return 0
}

type Command struct {
	place *game.Place
	actor *game.Actor
}

func (c Command) Execute(args []string) string {
	var itemName string
	if len(args) == 0 {
		return (*c.place).Description()
	} else {
		itemName = args[len(args)-1]
	}
	for _, object := range (*c.place).Objects() {
		if (*object).FriendlyName() == itemName {
			return (*object).Description()
		}
		if onItems, ok := (*object).ObjectsMap()[language.ON]; (*object).IsContainer() && ok {
			for _, innerItem := range onItems {
				if (*innerItem).FriendlyName() == itemName {
					return (*innerItem).Description()
				}
			}
		}
	}
	for _, object := range (*c.actor).Objects() {
		if (*object).FriendlyName() == itemName {
			return (*object).Description()
		}
	}
	return NothingFound
}
