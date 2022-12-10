package game

import (
	"errors"
	"fmt"

	"github.com/halpdesk/randomwalker/language"
	"github.com/halpdesk/randomwalker/screen"
)

var ErrPrepositionDoesNotExist = errors.New("preposition does not exist")
var ErrNotAContainer = errors.New("the item is not a container")
var ErrIsContainer = errors.New("the item is byitself a container")

type ItemType string

const FURNITURE ItemType = "furniture"
const CONTAINER ItemType = "container"
const ITEM ItemType = "item"

type Item struct {
	name        string
	description string
	itemType    ItemType
	article     language.Article
	attachments []language.Preposition
	items       map[language.Preposition][]*Item
}

func (i Item) FriendlyName() string {
	return i.name
}

func (i Item) Description() string {
	return screen.Color(fmt.Sprintf("%s %s", i.article, i.name), screen.RedFg, screen.BlackBg)
}

func (i *Item) PossibleAttachments() []language.Preposition {
	return i.attachments
}

func (i *Item) ItemMap() map[language.Preposition][]*Item {
	return i.items
}

func (i *Item) Items() []*Item {
	var items []*Item
	for position, _ := range i.items {
		for _, item := range i.items[position] {
			items = append(items, item)
		}
	}
	return items
}

func (i *Item) IsContainer() bool {
	return i.itemType == CONTAINER || i.itemType == FURNITURE
}

func (i *Item) PutItem(item *Item, preposition language.Preposition) error {
	if i.itemType != CONTAINER {
		return ErrNotAContainer
	}
	if item.itemType != CONTAINER {
		return ErrIsContainer
	}
	var exists bool
	for _, pre := range i.attachments {
		if pre == preposition {
			exists = true
		}
	}
	if !exists {
		return ErrPrepositionDoesNotExist
	}
	i.items[preposition] = append(i.items[preposition], item)
	return nil
}
