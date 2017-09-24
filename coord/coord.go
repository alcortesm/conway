package coord

type Coord struct {
	x, y uint
}

func New(x, y uint) *Coord {
	return &Coord{x: x, y: y}
}

func (c *Coord) X() uint {
	return c.x
}

func (c *Coord) Y() uint {
	return c.y
}
