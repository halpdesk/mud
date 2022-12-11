package item

import (
	"fmt"
	"strings"

	"github.com/halpdesk/mud/core/language"
	"github.com/halpdesk/mud/core/screen"
	"github.com/halpdesk/mud/game"
)

func New(name string, description string, article language.Article) game.Object {
	return &Item{
		name:        name,
		description: description,
		objectType:  game.ITEM,
		article:     article,
		attachments: nil,
		objects:     map[language.Preposition][]game.Object{},
	}
}

type Item struct {
	name        string
	description string
	objectType  game.ObjectType
	article     language.Article
	attachments []language.Preposition
	objects     map[language.Preposition][]game.Object
}

func (i *Item) FriendlyName() string {
	return i.name
}

func (i *Item) CursoryDescription() string {
	return screen.Color(fmt.Sprintf("%s %s", i.article, i.name), screen.RedFg, screen.BlackBg)
}

func (i *Item) Description() string {
	if len(i.objects) > 0 {
		return fmt.Sprintf("%s. %s", i.description, i.objectsDescription())
	}
	return fmt.Sprintf("%s", i.description)
}

func (i *Item) objectsDescription() string {
	prepositionDescriptions := []string{}
	for preposition := range i.objects {
		itemList := []string{}
		for _, item := range i.objects[preposition] {
			itemList = append(itemList, item.CursoryDescription())
		}
		prepositionDescriptions = append(prepositionDescriptions, fmt.Sprintf("There %s %s %s it", language.NumerusArticle(len(i.objects[preposition])), strings.Join(itemList, ", "), preposition))
	}
	return fmt.Sprintf("%s", strings.Join(prepositionDescriptions, ", "))
}

func (i *Item) PossibleAttachments() []language.Preposition {
	return i.attachments
}

func (i *Item) ObjectsMap() map[language.Preposition][]*game.Object {
	return nil
}

func (i *Item) Objects() []*game.Object {
	return nil
}

func (i *Item) RemoveObject(object *game.Object) error {
	return game.ErrIsContainer
}

func (i *Item) IsContainer() bool {
	return false
}

func (i *Item) PutObject(object *game.Object, preposition language.Preposition) error {
	return game.ErrNotAContainer
}
