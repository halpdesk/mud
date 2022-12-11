package room

import (
	"fmt"
	"strings"

	"github.com/halpdesk/mud/components/abstracts/container"
	"github.com/halpdesk/mud/core/language"
	"github.com/halpdesk/mud/core/screen"
	"github.com/halpdesk/mud/game"
)

func New(name string, description string, longDescription string, coordinates game.Coordinates) game.Place {
	cont := container.New([]language.Preposition{language.IN})
	return &Room{
		container:       &cont,
		name:            name,
		description:     description,
		longDescription: longDescription,
		coordinates:     coordinates,
		placeType:       game.ROOM,
	}
}

type Room struct {
	container       *game.Container
	name            string
	description     string
	longDescription string
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
	numberOfObjects := len((*r.container).Objects())
	if numberOfObjects == 0 {
		return ""
	}

	itemList := []string{}
	for _, item := range (*r.container).Objects() {
		itemList = append(itemList, (*item).CursoryDescription())
	}
	return fmt.Sprintf("In the %s there %s %s", r.placeType, language.NumerusArticle(numberOfObjects), strings.Join(itemList, ", "))
}

func (r *Room) Coordinates() game.Coordinates {
	return r.coordinates
}

func (r *Room) Container() *game.Container {
	return r.container
}
