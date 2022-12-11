package player

import (
	"github.com/halpdesk/mud/components/container"
	"github.com/halpdesk/mud/core/language"
	"github.com/halpdesk/mud/game"
)

func New(name string, hp int, startRoom *game.Place) game.Actor {
	bag := container.New(
		"Bag",
		"Made of leather. By pulling or loosing a thread, it is possible to close and open it",
		language.A,
		[]language.Preposition{language.IN},
	)
	return &Player{
		name:         name,
		healthPoints: hp,
		place:        startRoom,
		items:        []*game.Object{&bag},
	}
}

type Player struct {
	name         string
	healthPoints int
	place        *game.Place
	items        []*game.Object
}

func (p *Player) Name() string {
	return p.name
}

func (p *Player) Place() *game.Place {
	return p.place
}

func (p *Player) Objects() []*game.Object {
	var items []*game.Object
	for _, outerItem := range p.items {
		items = append(items, outerItem)
		if (*outerItem).IsContainer() {
			for _, innerItem := range (*outerItem).Objects() {
				items = append(items, innerItem)
			}
		}
	}
	return items
}

func (p *Player) GiveObject(object *game.Object) {
	p.items = append(p.items, object)
}

func (p *Player) Inventory() map[string][]*game.Object {
	inventory := make(map[string][]*game.Object, 0)
	inventory["self"] = []*game.Object{}
	for _, outerItem := range p.items {
		if (*outerItem).IsContainer() {

			inventory[(*outerItem).FriendlyName()] = []*game.Object{}
			for _, innerItem := range (*outerItem).Objects() {
				inventory[(*outerItem).FriendlyName()] = append(inventory[(*outerItem).FriendlyName()], innerItem)
			}
		} else {
			inventory["self"] = append(inventory["self"], outerItem)
		}
	}
	return inventory
}

func (p *Player) WalkToPlace(place *game.Place) {
	p.place = place
}
