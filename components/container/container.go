package container

import (
	"fmt"
	"strings"

	"github.com/halpdesk/mud/core/language"
	"github.com/halpdesk/mud/core/screen"
	"github.com/halpdesk/mud/game"
)

func New(name string, description string, article language.Article, attachments []language.Preposition) game.Object {
	return &Container{
		name:        name,
		description: description,
		objectType:  game.CONTAINER,
		article:     article,
		attachments: attachments,
		objects:     make(map[language.Preposition][]*game.Object, 0),
	}
}

type Container struct {
	name        string
	description string
	objectType  game.ObjectType
	article     language.Article
	attachments []language.Preposition
	objects     map[language.Preposition][]*game.Object
}

func (c *Container) FriendlyName() string {
	return c.name
}

func (c *Container) CursoryDescription() string {
	return screen.Color(fmt.Sprintf("%s %s", c.article, c.name), screen.RedFg, screen.BlackBg)
}

func (c *Container) Description() string {
	if len(c.objects) > 0 {
		return fmt.Sprintf("%s. %s", c.description, c.objectsDescription())
	}
	return fmt.Sprintf("%s", c.description)
}

func (c *Container) objectsDescription() string {
	prepositionDescriptions := []string{}
	for preposition := range c.objects {
		ContainerList := []string{}
		for _, object := range c.objects[preposition] {
			ContainerList = append(ContainerList, (*object).CursoryDescription())
		}
		prepositionDescriptions = append(prepositionDescriptions, fmt.Sprintf("There %s %s %s it", language.NumerusArticle(len(c.objects[preposition])), strings.Join(ContainerList, ", "), preposition))
	}
	return fmt.Sprintf("%s", strings.Join(prepositionDescriptions, ", "))
}

func (c *Container) PossibleAttachments() []language.Preposition {
	return c.attachments
}

func (c *Container) ObjectsMap() map[language.Preposition][]*game.Object {
	return c.objects
}

func (c *Container) Objects() []*game.Object {
	var objects []*game.Object
	for position := range c.objects {
		for _, object := range c.objects[position] {
			objects = append(objects, object)
		}
	}
	return objects
}

func (c *Container) IsContainer() bool {
	return true
}

func (c *Container) PutObject(object *game.Object, preposition language.Preposition) error {
	if !c.IsContainer() {
		return game.ErrNotAContainer
	}
	if (*object).IsContainer() {
		return game.ErrIsContainer
	}
	var exists bool
	for _, pre := range c.attachments {
		if pre == preposition {
			exists = true
		}
	}
	if !exists {
		return game.ErrPrepositionDoesNotExist
	}
	c.objects[preposition] = append(c.objects[preposition], object)
	return nil
}

func (c *Container) RemoveObject(object *game.Object) error {
	for preposition := range c.objects {
		for i, innerObject := range c.objects[preposition] {
			if object == innerObject {
				c.objects[preposition] = append(c.objects[preposition][:i], c.objects[preposition][i+1:]...)
				return nil
			}
		}
	}
	return game.ErrObjectDoesNotExistInContainer
}
