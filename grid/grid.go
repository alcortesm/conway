package grid

import (
	"fmt"

	"github.com/alcortesm/conway/conway"
	"github.com/alcortesm/conway/coord"
)

// Grid represents a snapshot of a universe.
// The zero value of this type is not safe, use the function New below.
// The minimum width and height for a universe is 3 cells each.
type Grid struct {
	width, height int
	cells         []bool // true means alive
}

const (
	// MinWidth is the minimum width for a grid.
	MinWidth = 3
	// MinHeight is the minimum height for a grid.
	MinHeight = 3
)

// New creates a new grid with the given width and height (number of cells)
// and the given list of alive cells.
// The grid has its origin of coordinates at the upper left corner.
// Returns an error if the width or the height are smaller than MinWidth and
// MinHeight or if any of the alive cells are out of bounds.
func New(width, height int, alives []conway.Coord) (*Grid, error) {
	if width < MinWidth {
		return nil, fmt.Errorf("width must be >= than %d, was %d",
			MinWidth, width)
	}
	if height < MinHeight {
		return nil, fmt.Errorf("height must be >= than %d, was %d",
			MinHeight, height)
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

// NewRandom returns a grid with n random alive cells inside a grid with the given dimensions.
func NewRandom(w, h, n int) *Grid {
	alives := make([]conway.Coord, n)
	for i := 0; i < n; i++ {
		alives[i] = coord.NewRandom(w, h)
	}
	g, err := New(w, h, alives)
	if err != nil {
		panic(fmt.Sprintf("cannot create grid: %v", err))
	}
	return g
}

// Index returns the index in the internal slice of a grid for the given c.
// Cells will be located in the array by rows in ordinate increasing order:
//  a b
//  c d   ->  a b c d e f
//  e f
func (g *Grid) index(c conway.Coord) (int, error) {
	return c.Y()*g.Width() + c.X(), g.checkBounds(c)
}

func (g *Grid) checkBounds(c conway.Coord) error {
	x, y := c.X(), c.Y()
	w, h := g.Width(), g.Height()
	if x < 0 || x >= w {
		return fmt.Errorf("invalid abscissa value (%d) for grid (width = %d)",
			x, w)
	}
	if y < 0 || y >= h {
		return fmt.Errorf("invalid ordinate value (%d) for grid (height = %d)",
			y, h)
	}
	return nil
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
