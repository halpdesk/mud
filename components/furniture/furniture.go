package furniture

import (
	"fmt"

	"github.com/halpdesk/mud/components/abstracts/container"
	"github.com/halpdesk/mud/core/language"
	"github.com/halpdesk/mud/core/screen"
	"github.com/halpdesk/mud/game"
)

func New(name string, description string, article language.Article, attachments []language.Preposition) game.Object {
	container := container.New(attachments)
	return &Furniture{
		container:   &container,
		name:        name,
		description: description,
		article:     article,
	}
}

type Furniture struct {
	container   *game.Container
	name        string
	description string
	article     language.Article
}

func (f *Furniture) ObjectType() game.ObjectType {
	return game.FURNITURE
}

func (f *Furniture) FriendlyName() string {
	return f.name
}

func (f *Furniture) CursoryDescription() string {
	return fmt.Sprintf("%s %s", f.article, screen.Color(f.name, screen.RedFg, screen.BlackBg))
}

func (f *Furniture) Description() string {
	if len((*f.container).Objects()) > 0 {
		return fmt.Sprintf("%s. %s", f.description, (*f.container).ObjectsDescription())
	}
	return fmt.Sprintf("%s", f.description)
}

func (f *Furniture) IsContainer() bool {
	return true
}

func (f *Furniture) Container() *game.Container {
	return f.container
}
