package grid

import (
	"fmt"

	"github.com/alcortesm/conway/conway"
)

// Grid represents a snapshot of a universe.
// The zero value of this type is not safe, use the function New below.
// The minimum width and height for a universe is 3 cells each.
type Grid struct {
	width, height uint
	cells         []bool // true means alive
}

const (
	minWidth  = 3
	minHeight = 3
)

// New creates a new grid with the given width and height (number of cells)
// and the given list of alive cells.
// Returns an error if the width or the height is smaller than 3 or if any
// of the alive cells are out of bounds.
func New(width, height uint, alives []conway.Coord) (*Grid, error) {
	if width < minWidth {
		return nil, fmt.Errorf("width must be >= than %d, was %d",
			minWidth, width)
	}
	if height < minHeight {
		return nil, fmt.Errorf("height must be >= than %d, was %d",
			minHeight, height)
	}
	g := &Grid{
		width:  width,
		height: height,
		cells:  make([]bool, width*height),
	}
	for i, c := range alives {
		p, err := g.index(c)
		if err != nil {
			return nil, fmt.Errorf("alive #%d: %v", i, err)
		}
		g.cells[p] = true
	}
	return g, nil
}

// Index returns the index in the internal slice of a grid for the given c.
// Cells will be located in the array by rows in ordinate increasing order:
//  a b
//  c d   ->  a b c d e f
//  e f
func (g *Grid) index(c conway.Coord) (int, error) {
	return int(c.Y()*g.Width() + c.X()), g.checkBounds(c)
}

func (g *Grid) checkBounds(c conway.Coord) error {
	x, y := c.X(), c.Y()
	w, h := g.Width(), g.Height()
	if x >= w {
		return fmt.Errorf("abscissa value too high (%d) for grid (width = %d)",
			x, w)
	}
	if y >= h {
		return fmt.Errorf("ordinate value too high (%d) for grid (height = %d)",
			y, h)
	}
	return nil
}

// Width returns the width of the universe (number of cells).
// Implements conway.Grid.
func (g *Grid) Width() uint {
	return g.width
}

// Height returns the height of the universe (number of cells).
// Implements conway.Grid.
func (g *Grid) Height() uint {
	return g.height
}

// IsAlive returns if the cell at coordinates c is alive.  Returns
// an error if c is out of bounds.
// Implements conway.Grid.
func (g *Grid) IsAlive(c conway.Coord) (bool, error) {
	i, err := g.index(c)
	if err != nil {
		return false, err
	}
	return g.cells[i], nil
}
