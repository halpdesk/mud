package game

import (
	"errors"

	"github.com/halpdesk/mud/core/language"
)

type Coordinates interface {
	GetY() int
	GetX() int
	GetCoordinates() (int, int)
}

type World interface {
	Places() []*Place
	AvailableCoordinates() []Coordinates
}

type Place interface {
	FriendlyName() string
	CursoryDescription() string
	Description() string
	Objects() []*Object
	Coordinates() Coordinates
}

type PlaceType string

const ROOM PlaceType = "room"
const AREA PlaceType = "area"

type Object interface {
	FriendlyName() string
	CursoryDescription() string
	Description() string
	// itemsDescription() string
	PossibleAttachments() []language.Preposition
	ObjectsMap() map[language.Preposition][]*Object
	Objects() []*Object
	IsContainer() bool
	PutObject(object *Object, preposition language.Preposition) error
	RemoveObject(object *Object) error
}

// ContainerItem
// Container

var ErrPrepositionDoesNotExist = errors.New("preposition does not exist")
var ErrObjectDoesNotExistInContainer = errors.New("the object does not exist in this container")
var ErrNotAContainer = errors.New("the item is not a container")
var ErrIsContainer = errors.New("the item is byitself a container")

type ObjectType string

const FURNITURE ObjectType = "furniture"
const CONTAINER ObjectType = "container"
const ITEM ObjectType = "item"

type Actor interface {
	Name() string
	Place() *Place
	Objects() []*Object
	GiveObject(object *Object)
	Inventory() map[string][]*Object
	WalkToPlace(r *Place)
}
