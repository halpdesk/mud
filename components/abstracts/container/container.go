package container

import (
	"fmt"
	"strings"

	"github.com/halpdesk/mud/core/language"
	"github.com/halpdesk/mud/game"
)

func New(attachments []language.Preposition) game.Container {
	objectMap := make(map[language.Preposition][]*game.Object, 0)
	return &Container{
		attachments: attachments,
		objects:     objectMap,
	}
}

type Container struct {
	attachments []language.Preposition
	objects     map[language.Preposition][]*game.Object
}

func (c *Container) ObjectsDescription() string {
	prepositionDescriptions := []string{}
	for preposition := range c.objects {
		objectsList := []string{}
		for _, object := range c.objects[preposition] {
			objectsList = append(objectsList, (*object).CursoryDescription())
		}
		prepositionDescriptions = append(prepositionDescriptions, fmt.Sprintf("There %s %s %s it", language.NumerusArticle(len(c.objects[preposition])), strings.Join(objectsList, ", "), preposition))
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
	for _, o := range c.objects {
		objects = append(objects, o...)
	}
	return objects
}

func (c *Container) PutObject(object *game.Object, preposition language.Preposition) error {
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
			if *object == *innerObject {
				c.objects[preposition] = append(c.objects[preposition][:i], c.objects[preposition][i+1:]...)
				return nil
			}
		}
	}
	return game.ErrObjectDoesNotExistInContainer
}
