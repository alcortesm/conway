/*
Package conway defines interfaces to generate animations of the Conway's game of life.
*/
package conway

import (
	"io"

	"github.com/alcortesm/conway/coord"
)

// Universe is a collection of cells that evolves over time.
type Universe interface {
	// Status returns the current status of the universe as a grid.
	Status() Grid
	// Tick makes the universe evolve a single round (a tick).
	Tick() error
}

// Animator represents a collection of grids that can be rendered in a graphical format
// and stored in a file.
type Animator interface {
	// Add adds a grid to the collection to be used as a photogram in the animate method.
	Add(Grid) error
	// Encode creates an animation of all the added photograms and store it in
	// the given writer.
	Encode(w io.Writer) error
}

// Grid represents an inmutable snapshot of a universe.
type Grid interface {
	// Width returns the width of the universe (number of cells).
	Width() int
	// Height returns the height of the universe (number of cells).
	Height() int
	// IsAlive returns if the cell at coordinates x, y is alive.
	// Returns an error if x or y are out of bounds.
	IsAlive(Coord) (bool, error)
}

// Coord represents the position of a cell in a grid.
type Coord interface {
	// X returns the cell's abscissa.
	X() int
	// Y returns the cell's ordinate.
	Y() int
}

// CoordEqual returns if two coordinates refer to the same cell.
func CoordEqual(a, b Coord) bool {
	return a.X() == b.X() && a.Y() == b.Y()
}

// Evolve returns the alive cells of a grid in its next tick.
func Evolve(old Grid) ([]Coord, error) {
	ret := []Coord{}
	for x := 0; x < old.Width(); x++ {
		for y := 0; y < old.Height(); y++ {
			current := coord.New(x, y)
			a, err := old.IsAlive(current)
			if err != nil {
				return nil, err
			}
			c := countNeighbours(old, x, y)
			if laws(a, c) {
				ret = append(ret, current)
			}
		}
	}
	return ret, nil
}

func countNeighbours(g Grid, x, y int) int {
	return 0
}

func laws(alive bool, neighbours int) bool {
	return false
}
