package game

import (
	"fmt"
	"strings"

	"github.com/halpdesk/mud/language"
	"github.com/halpdesk/mud/screen"
)

type RoomType string

const ROOM RoomType = "room"
const AREA RoomType = "area"

type Room struct {
	name        string
	description string
	items       []*Item
	coordinates Coordinates
	roomType    RoomType
}

func (r *Room) FriendlyName() string {
	return r.name
}

func (r *Room) CursoryDescription() string {
	return r.name
}

func (r *Room) Description() string {
	name := screen.Color(r.name, screen.LightBlueFg, screen.BlackBg)
	description := screen.Color(r.description, screen.BlueFg, screen.BlackBg)
	coordinates := screen.Color(fmt.Sprintf("[%d, %d]", r.coordinates.X, r.coordinates.Y), screen.LightGrayFg, screen.BlackBg)
	return fmt.Sprintf("%s %s %s\n%s", name, description, coordinates, r.itemsDescription())
}

func (r *Room) itemsDescription() string {
	itemList := []string{}
	for _, item := range r.items {
		itemList = append(itemList, item.Description())
	}
	return fmt.Sprintf("In the %s there %s %s.", r.roomType, r.itemNumerusArticle(), strings.Join(itemList, ", "))
}

func (r *Room) itemNumerusArticle() string {
	numerusArticle := "are"
	if len(r.items) < 2 {
		numerusArticle = "is"
	}
	if len(r.items) < 1 {
		numerusArticle = "nothing"
	}
	return numerusArticle
}

func (r *Room) Coordinates() Coordinates {
	return r.coordinates
}

var startRoom = Room{
	name:        "Start",
	roomType:    ROOM,
	description: "This is the start",
	items: []*Item{{
		name:        "Desk",
		description: "Wooden and old. Has int stains on it.",
		itemType:    CONTAINER,
		article:     language.A,
		attachments: []language.Preposition{language.ON},
		items: map[language.Preposition][]*Item{language.ON: {{
			name:        "Clock",
			description: "Regular clock. It ticks",
			itemType:    ITEM,
			article:     language.A,
			attachments: []language.Preposition{},
			items:       map[language.Preposition][]*Item{},
		}}},
	}},
	coordinates: Coordinates{
		X: 0,
		Y: 0,
	},
}

var rooms = []*Room{
	&startRoom,
	{
		name:        "End",
		roomType:    ROOM,
		description: "This is the end",
		items: []*Item{{
			name:    "Mirror",
			article: language.A,
		}},
		coordinates: Coordinates{
			X: 1,
			Y: 0,
		},
	},
}
