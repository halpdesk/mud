package item

import (
	"fmt"

	"github.com/halpdesk/mud/core/language"
	"github.com/halpdesk/mud/core/screen"
	"github.com/halpdesk/mud/game"
)

func New(name string, description string, article language.Article) game.Object {
	return &Item{
		name:        name,
		description: description,
		article:     article,
	}
}

type Item struct {
	name        string
	description string
	article     language.Article
}

func (i *Item) ObjectType() game.ObjectType {
	return game.ITEM
}

func (i *Item) FriendlyName() string {
	return i.name
}

func (i *Item) CursoryDescription() string {
	return fmt.Sprintf("%s %s", i.article, screen.Color(i.name, screen.RedFg, screen.BlackBg))
}

func (i *Item) Description() string {
	return fmt.Sprintf("%s", i.description)
}

func (i *Item) IsContainer() bool {
	return false
}

func (i *Item) Container() *game.Container {
	return nil
}
