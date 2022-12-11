package take

import (
	"fmt"

	"github.com/halpdesk/mud/core/language"
	"github.com/halpdesk/mud/core/screen"
	"github.com/halpdesk/mud/game"
)

var NothingFound = "There is no such thing here"
var AlreadyHave = "You already have that"
var CannotPickUpFurniture = "You cannot pick up anything of that size"

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
	itemName := args[0]
	var item game.Object
	var container *game.Container
	for _, object := range (*(*c.place).Container()).Objects() {
		if language.Equal((*object).FriendlyName(), itemName) {
			if (*object).ObjectType() == game.FURNITURE {
				return CannotPickUpFurniture
			}
			item = *object
			container = (*c.place).Container()
		}
		if (*object).IsContainer() {
			for _, preposition := range (*(*object).Container()).PossibleAttachments() {
				if onItems, ok := (*(*object).Container()).ObjectsMap()[preposition]; ok {
					for _, innerItem := range onItems {
						if language.Equal((*innerItem).FriendlyName(), itemName) {
							container = (*object).Container()
							item = *innerItem
						}
					}
				}
			}
		}
	}
	for _, object := range (*c.actor).Objects() {
		if language.Equal((*object).FriendlyName(), itemName) {
			return AlreadyHave
		}
	}
	if item != nil && container != nil {
		(*c.actor).GiveObject(&item)
		(*container).RemoveObject(&item)
		return fmt.Sprintf("You picked up the %s", screen.Color(item.FriendlyName(), screen.RedFg, screen.BlackBg))
	}
	return NothingFound
}
