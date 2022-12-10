package game

import "github.com/halpdesk/mud/language"

func NewPlayer(name string, hp int) Player {
	return Player{
		name:         name,
		healthPoints: hp,
		room:         &startRoom,
		items: []*Item{{
			name:        "Bag",
			description: "Made of leather. By pulling or loosing a thread, it is possible to close and open it",
			article:     language.A,
			attachments: []language.Preposition{language.IN},
			items:       map[language.Preposition][]*Item{},
		}},
	}
}

type Player struct {
	name         string
	healthPoints int
	room         *Room
	items        []*Item
	// containers    []*Container
}

func (p *Player) Name() string {
	return p.name
}

func (p *Player) Room() *Room {
	return p.room
}

func (p *Player) Inventory() map[string][]*Item {
	inventory := make(map[string][]*Item, 0)
	inventory["self"] = []*Item{}
	for _, outerItem := range p.items {
		if outerItem.IsContainer() {

			inventory[outerItem.FriendlyName()] = []*Item{}
			for _, innerItem := range outerItem.Items() {
				inventory[outerItem.FriendlyName()] = append(inventory[outerItem.FriendlyName()], innerItem)
			}
		} else {
			inventory["self"] = append(inventory["self"], outerItem)
		}
	}
	return inventory
}

func (p *Player) WalkToRoom(r *Room) {
	p.room = r
}
