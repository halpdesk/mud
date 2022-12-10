package look

import (
	"github.com/halpdesk/mud/game"
	"github.com/halpdesk/mud/language"
)

var NothingFound = "There is no such thing here"

func New(player *game.Player, room *game.Room) Command {
	return Command{
		room:   room,
		player: player,
	}
}

func (c Command) Arity() int {
	return 0
}

type Command struct {
	room   *game.Room
	player *game.Player
}

func (c Command) Execute(args []string) string {
	var itemName string
	if len(args) == 0 {
		return c.room.Description()
	} else {
		itemName = args[len(args)-1]
	}
	for _, item := range c.room.Items() {
		if item.FriendlyName() == itemName {
			return item.Description()
		}
		if onItems, ok := item.ItemMap()[language.ON]; item.IsContainer() && ok {
			for _, innerItem := range onItems {
				if innerItem.FriendlyName() == itemName {
					return innerItem.Description()
				}
			}
		}
	}
	for _, item := range c.player.Items() {
		if item.FriendlyName() == itemName {
			return item.Description()
		}
	}
	return NothingFound
}
