package grid

import "fmt"

// Grid represents a snapshot of a universe.
// The zero value of this type is not safe, use the function New below.
// The minimum width and height for a universe is 3 cells each.
type Grid struct {
	width, height int
}

const (
	minWidth  = 3
	minHeight = 3
)

// New creates a new grid with the given width and height (number of cells).
// Returns an error if the width or the height is smaller than 3.
func New(width, height int) (*Grid, error) {
	if width < minWidth {
		return nil, fmt.Errorf("width must be >= than %d, was %d",
			minWidth, width)
	}
	if height < minHeight {
		return nil, fmt.Errorf("height must be >= than %d, was %d",
			minHeight, height)
	}
	return &Grid{
		width:  width,
		height: height,
	}, nil
}

// Width returns the width of the universe (number of cells).
// Implements conway.Grid.
func (g *Grid) Width() int {
	return g.width
}

// Height returns the height of the universe (number of cells).
// Implements conway.Grid.
func (g *Grid) Height() int {
	return g.height
}

// Get returns if the cell at coordinates x, y is alive.  Returns
// an error if x or y are out of bounds.
// Implements conway.Grid.
func (g *Grid) Get(x, y int) (bool, error) {
	return false, fmt.Errorf("TODO")
}

// Set sets the cell at coordinates x, y to value v (true for alive,
// false for dead).  Returns an error if x or y are out of bounds.
// Implements conway.Grid.
func (g *Grid) Set(x, y int, v bool) error {
	return fmt.Errorf("TODO")
}
