package coordinates

import "github.com/halpdesk/mud/game"

func New(x, y int) game.Coordinates {
	return Coordinates{
		x: x,
		y: y,
	}
}

type Coordinates struct {
	x int
	y int
}

func (c Coordinates) GetCoordinates() (int, int) {
	return c.x, c.y
}

func (c Coordinates) GetX() int {
	return c.x
}

func (c Coordinates) GetY() int {
	return c.y
}
