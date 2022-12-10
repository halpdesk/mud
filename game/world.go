package game

func NewWorld() World {
	return World{
		rooms: rooms,
	}
}

type World struct {
	rooms []*Room
}

func (w World) GetRooms() []*Room {
	return w.rooms
}

func (w World) GetRoomCoordinates() []Coordinates {
	var coordinates []Coordinates
	for _, room := range w.rooms {
		coordinates = append(coordinates, room.coordinates)
	}
	return coordinates
}
