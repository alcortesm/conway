package random

import (
	"fmt"

	"github.com/alcortesm/conway/conway"
	"github.com/alcortesm/conway/grid"
)

// Random is a Conway universe that starts with random alives.
type Random struct{}

// New returns a new Random universe.
func New() *Random {
	return &Random{}
}

// Status implements conway.Universe.
func (r *Random) Status() conway.Grid {
	g, err := grid.New(5, 10, []conway.Coord{})
	if err != nil {
		panic(fmt.Sprintf("cannot create grid: %v", err))
	}
	return g
}

// Tick implements conway.Universe.
func (r *Random) Tick() {}
