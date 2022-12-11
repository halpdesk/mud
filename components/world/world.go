package world

import "github.com/halpdesk/mud/game"

func New(places []*game.Place) game.World {
	return &World{
		places: places,
	}
}

type World struct {
	places []*game.Place
}

func (w *World) Places() []*game.Place {
	return w.places
}

func (w *World) AvailableCoordinates() []game.Coordinates {
	var coordinates []game.Coordinates
	for _, place := range w.places {
		p := *place
		coordinates = append(coordinates, p.Coordinates())
	}
	return coordinates
}
