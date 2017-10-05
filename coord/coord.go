package coord

// Coord represents a cell in a grid.
type Coord struct {
	x, y uint
}

// New creates a new coord, representing a cell in a grid in the given
// position.  Positions are represented as cartesian coordinates, where
// the origin is at the upper left of the grid.
func New(x, y uint) *Coord {
	return &Coord{x: x, y: y}
}

// X returns the cell's abscissa.
// Implements conway.Coord.
func (c *Coord) X() uint {
	return c.x
}

// Y returns the cell's ordinate.
// Implements conway.Coord.
func (c *Coord) Y() uint {
	return c.y
}
