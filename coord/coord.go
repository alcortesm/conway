package coord

import "math/rand"

// Coord represents a cell in a grid.
type Coord struct {
	x, y int
}

// New creates a new coord, representing a cell in a grid in the given
// position.  Positions are represented as cartesian coordinates.
func New(x, y int) *Coord {
	return &Coord{x: x, y: y}
}

// NewRandom creates a new coord in a random position inside a grid of the given dimensions.
func NewRandom(w, h int) *Coord {
	x := rand.Intn(w)
	y := rand.Intn(h)
	return New(x, y)
}

// X returns the cell's abscissa.
// Implements conway.Coord.
func (c *Coord) X() int {
	return c.x
}

// Y returns the cell's ordinate.
// Implements conway.Coord.
func (c *Coord) Y() int {
	return c.y
}
