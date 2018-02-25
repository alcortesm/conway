package evolve

import (
	"github.com/alcortesm/conway/conway"
	"github.com/alcortesm/conway/coord"
)

// Evolve returns the alive cells of a grid in its next tick.
func Evolve(old conway.Grid) ([]conway.Coord, error) {
	ret := []conway.Coord{}
	for x := 0; x < old.Width(); x++ {
		for y := 0; y < old.Height(); y++ {
			current := coord.New(x, y)
			a, err := old.IsAlive(current)
			if err != nil {
				return nil, err
			}
			c := countNeighbours(old, current)
			if laws(a, c) {
				ret = append(ret, current)
			}
		}
	}
	return ret, nil
}

func countNeighbours(g conway.Grid, c conway.Coord) int {
	sum := 0
	minX := max(c.X()-1, 0)
	minY := max(c.Y()-1, 0)
	maxX := min(c.X()+1, g.Width()-1)
	maxY := min(c.Y()+1, g.Height()-1)
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			if x == c.X() && y == c.Y() {
				continue
			}
			alive, err := g.IsAlive(coord.New(x, y))
			if err != nil {
				panic(err)
			}
			if alive {
				sum++
			}
		}
	}
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func laws(alive bool, neighbours int) bool {
	if alive && neighbours == 2 {
		return true
	}
	if neighbours == 3 {
		return true
	}
	return false
}
