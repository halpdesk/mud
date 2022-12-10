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
	name            string
	description     string
	longDescription string
	items           []*Item
	coordinates     Coordinates
	roomType        RoomType
}

func (r *Room) FriendlyName() string {
	return r.name
}

func (r *Room) CursoryDescription() string {
	name := screen.Color(r.name, screen.LightBlueFg, screen.BlackBg)
	description := screen.Color(r.description, screen.BlueFg, screen.BlackBg)
	coordinates := screen.Color(fmt.Sprintf("[%d, %d]", r.coordinates.X, r.coordinates.Y), screen.LightGrayFg, screen.BlackBg)
	return fmt.Sprintf("%s %s\n%s. %s", name, coordinates, description, r.itemsDescription())
}

func (r *Room) Description() string {
	name := screen.Color(r.name, screen.LightBlueFg, screen.BlackBg)
	description := screen.Color(r.description, screen.BlueFg, screen.BlackBg)
	coordinates := screen.Color(fmt.Sprintf("[%d, %d]", r.coordinates.X, r.coordinates.Y), screen.LightGrayFg, screen.BlackBg)
	return fmt.Sprintf("%s %s\n%s. %s. %s", name, coordinates, description, r.longDescription, r.itemsDescription())
}

func (r *Room) itemsDescription() string {
	itemList := []string{}
	for _, item := range r.items {
		itemList = append(itemList, item.CursoryDescription())
	}
	return fmt.Sprintf("In the %s there %s %s", r.roomType, language.ItemNumerusArticle(len(r.items)), strings.Join(itemList, ", "))
}

func (r *Room) Items() []*Item {
	return r.items
}

func (r *Room) Coordinates() Coordinates {
	return r.coordinates
}

var startRoom = Room{
	name:            "Start",
	roomType:        ROOM,
	description:     "This is the start",
	longDescription: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum",
	items: []*Item{{
		name:        "Desk",
		description: "It is wooden and old and has stains on it",
		itemType:    CONTAINER,
		article:     language.A,
		attachments: []language.Preposition{language.ON},
		items: map[language.Preposition][]*Item{language.ON: {{
			name:        "Clock",
			description: "It's a regular clock, tick tock",
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
		name:            "End",
		roomType:        ROOM,
		description:     "This is the end",
		longDescription: "Enim blandit volutpat maecenas volutpat blandit. A diam sollicitudin tempor id eu nisl nunc mi. Diam sollicitudin tempor id eu nisl. Nec tincidunt praesent semper feugiat nibh sed. Vulputate eu scelerisque felis imperdiet proin fermentum leo",
		items: []*Item{{
			name:        "Mirror",
			description: "There is you",
			article:     language.A,
		}},
		coordinates: Coordinates{
			X: 1,
			Y: 0,
		},
	},
}
