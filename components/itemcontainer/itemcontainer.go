package itemcontainer

import (
	"fmt"

	"github.com/halpdesk/mud/components/abstracts/container"
	"github.com/halpdesk/mud/core/language"
	"github.com/halpdesk/mud/core/screen"
	"github.com/halpdesk/mud/game"
)

func New(name string, description string, article language.Article, attachments []language.Preposition) game.Object {
	con := container.New(attachments)
	return &ItemContainer{
		container:   &con,
		name:        name,
		description: description,
		article:     article,
	}
}

type ItemContainer struct {
	container   *game.Container
	name        string
	description string
	article     language.Article
}

func (ic *ItemContainer) ObjectType() game.ObjectType {
	return game.ITEMCONTAINER
}

func (ic *ItemContainer) FriendlyName() string {
	return ic.name
}

func (ic *ItemContainer) CursoryDescription() string {
	return fmt.Sprintf("%s %s", ic.article, screen.Color(ic.name, screen.RedFg, screen.BlackBg))
}

func (ic *ItemContainer) Description() string {
	if len((*ic.container).Objects()) > 0 {
		return fmt.Sprintf("%s. %s", ic.description, (*ic.container).ObjectsDescription())
	}
	return fmt.Sprintf("%s", ic.description)
}

func (ic *ItemContainer) IsContainer() bool {
	return true
}

func (ic *ItemContainer) Container() *game.Container {
	return ic.container
}
