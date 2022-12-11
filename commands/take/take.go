package take

import (
	"github.com/halpdesk/mud/game"
)

var NothingFound = "There is no such thing here"
var AlreadyHave = "You already have that"

func New(actor *game.Actor, place *game.Place) Command {
	return Command{
		place: place,
		actor: actor,
	}
}

func (c Command) Arity() int {
	return 1
}

type Command struct {
	place *game.Place
	actor *game.Actor
}

func (c Command) Execute(args []string) string {
	// itemName := args[0]
	// var item game.Object
	// var container game.Object
	// for _, object := range (*c.place).Objects() {
	// 	if (*object).FriendlyName() == itemName && !(*object).IsFurniture() {
	// 		item = *object
	// 		container = *c.place
	// 	}
	// 	if onItems, ok := (*object).ObjectsMap()[language.ON]; (*object).IsContainer() && ok {
	// 		for _, innerItem := range onItems {
	// 			if (*innerItem).FriendlyName() == itemName {
	// 				container = *object
	// 				item = *innerItem
	// 			}
	// 		}
	// 	}
	// }
	// for _, object := range (*c.actor).Objects() {
	// 	if (*object).FriendlyName() == itemName {
	// 		return AlreadyHave
	// 	}
	// }
	// (*c.actor).GiveObject(&item)
	return NothingFound
}
