package room

import (
	"fmt"
	"strings"

	"github.com/halpdesk/mud/core/language"
	"github.com/halpdesk/mud/core/screen"
	"github.com/halpdesk/mud/game"
)

func New(name string, description string, longDescription string, objects []*game.Object, coordinates game.Coordinates) game.Place {
	return &Room{
		name:            name,
		description:     description,
		longDescription: longDescription,
		objects:         objects,
		coordinates:     coordinates,
		placeType:       game.ROOM,
	}
}

type Room struct {
	name            string
	description     string
	longDescription string
	objects         []*game.Object
	coordinates     game.Coordinates
	placeType       game.PlaceType
}

func (r *Room) FriendlyName() string {
	return r.name
}

func (r *Room) CursoryDescription() string {
	name := screen.Color(r.name, screen.LightBlueFg, screen.BlackBg)
	description := screen.Color(r.description, screen.BlueFg, screen.BlackBg)
	coordinates := screen.Color(fmt.Sprintf("[%d, %d]", r.coordinates.GetX(), r.coordinates.GetY()), screen.LightGrayFg, screen.BlackBg)
	return fmt.Sprintf("%s %s\n%s. %s", name, coordinates, description, r.objectsDescription())
}

func (r *Room) Description() string {
	name := screen.Color(r.name, screen.LightBlueFg, screen.BlackBg)
	description := screen.Color(r.description, screen.BlueFg, screen.BlackBg)
	coordinates := screen.Color(fmt.Sprintf("[%d, %d]", r.coordinates.GetX(), r.coordinates.GetY()), screen.LightGrayFg, screen.BlackBg)
	return fmt.Sprintf("%s %s\n%s. %s. %s", name, coordinates, description, r.longDescription, r.objectsDescription())
}

func (r *Room) objectsDescription() string {
	itemList := []string{}
	for _, item := range r.objects {
		itemList = append(itemList, (*item).CursoryDescription())
	}
	return fmt.Sprintf("In the %s there %s %s", r.placeType, language.NumerusArticle(len(r.objects)), strings.Join(itemList, ", "))
}

func (r *Room) Objects() []*game.Object {
	return r.objects
}

func (r *Room) Coordinates() game.Coordinates {
	return r.coordinates
}
