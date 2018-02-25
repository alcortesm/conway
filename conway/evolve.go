package conway

import "github.com/alcortesm/conway/coord"

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
	if alive && neighbours == 2 {
		return true
	}
	if neighbours == 3 {
		return true
	}
	return false
}
