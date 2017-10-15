package grid_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/alcortesm/conway/conway"
	"github.com/alcortesm/conway/coord"
	"github.com/alcortesm/conway/grid"
)

func TestSizeOK(t *testing.T) {
	testSizeOK(t, 3, 3)
	testSizeOK(t, 3, 4)
	testSizeOK(t, 4, 3)
	testSizeOK(t, 4, 4)
	testSizeOK(t, 10, 10)
	testSizeOK(t, 20, 10)
	testSizeOK(t, 10, 20)
	testSizeOK(t, 10000, 20000)
}

func testSizeOK(t *testing.T, width, height uint) {
	t.Helper()
	g, err := grid.New(width, height, nil)
	if err != nil {
		t.Fatalf("cannot create grid: %v", err)
	}
	if w := g.Width(); w != width {
		t.Errorf("wrong width: expected %d, got %d", width, w)
	}
	if h := g.Height(); h != height {
		t.Errorf("wrong height: expected %d, got %d", height, h)
	}
}

func TestSizeError(t *testing.T) {
	testSizeError(t, 2, 3)
	testSizeError(t, 3, 2)
	testSizeError(t, 2, 2)
	testSizeError(t, 0, 0)
}

func testSizeError(t *testing.T, width, height uint) {
	t.Helper()
	_, err := grid.New(width, height, nil)
	if err == nil {
		t.Errorf("new grid was supposed to fail and it did not")
	}
}

func TestIsAlive(t *testing.T) {
	for _, tt := range []struct {
		name   string
		w, h   uint
		alives []conway.Coord
	}{
		{"empty", 3, 3, nil},
		{"full", 3, 3, []conway.Coord{
			coord.New(0, 0),
			coord.New(0, 1),
			coord.New(0, 2),
			coord.New(1, 0),
			coord.New(1, 1),
			coord.New(1, 2),
			coord.New(2, 0),
			coord.New(2, 1),
			coord.New(2, 2),
		}},
		{"diagonal", 3, 3, []conway.Coord{
			coord.New(0, 0),
			coord.New(1, 1),
			coord.New(2, 2),
		}},
		{"non-diagonal", 3, 3, []conway.Coord{
			coord.New(0, 1),
			coord.New(0, 2),
			coord.New(1, 0),
			coord.New(1, 2),
			coord.New(2, 0),
			coord.New(2, 1),
		}},
		{"big", 20, 20, []conway.Coord{
			coord.New(0, 0),
			coord.New(7, 17),
			coord.New(17, 6),
			coord.New(19, 19),
		}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			testIsAlive(t, tt.w, tt.h, tt.alives)
		})
	}
}

func testIsAlive(t *testing.T, w, h uint, alives []conway.Coord) {
	g, err := grid.New(w, h, alives)
	if err != nil {
		t.Fatalf("cannot create grid: %v", err)
	}
	var x uint
	var y uint
	for x = 0; x < w; x++ {
		for y = 0; y < h; y++ {
			c := coord.New(x, y)
			expected := contains(alives, c)
			ia, err := g.IsAlive(c)
			if err != nil {
				t.Errorf("error calling IsAlive: %v", err)
			}
			if ia != expected {
				t.Errorf("wrong IsAlive output: (%d, %d) was %v, should be %v",
					c.X(), c.Y(), ia, expected)
			}
		}
	}
}

func contains(list []conway.Coord, c conway.Coord) bool {
	for _, e := range list {
		if conway.CoordEqual(e, c) {
			return true
		}
	}
	return false
}

func BenchmarkSize(b *testing.B) {
	for _, side := range []uint{300, 600, 900, 1200, 1500, 1800, 2100, 2400, 2700} {
		name := fmt.Sprintf("side=%d", side)
		all := allCoords(side)
		nAlives := percentOf(20, side*side)
		b.Run(name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				b.StopTimer()
				alives := randomSubSlice(nAlives, all)
				b.StartTimer()
				createGridAndCheckAllCells(b, side, alives)
			}
		})
	}
}

func allCoords(side uint) []conway.Coord {
	all := make([]conway.Coord, 0, side*side)
	var x, y uint
	for x = 0; x < side; x++ {
		for y = 0; y < side; y++ {
			all = append(all, coord.New(x, y))
		}
	}
	return all
}

func percentOf(p, total uint) uint {
	return total * p / 100
}

func randomSubSlice(n uint, all []conway.Coord) []conway.Coord {
	ret := make([]conway.Coord, n)
	randomIndexes := rand.Perm(len(all))
	for i, v := range randomIndexes[0:n] {
		ret[i] = all[v]
	}
	return ret
}

func createGridAndCheckAllCells(b *testing.B, side uint, alives []conway.Coord) {
	g, err := grid.New(side, side, alives)
	if err != nil {
		b.Fatalf("cannot create grid: %v", err)
	}
	var x, y uint
	for x = 0; x < side; x++ {
		for y = 0; y < side; y++ {
			if _, err := g.IsAlive(coord.New(x, y)); err != nil {
				b.Fatalf("checking if (%d, %d) is alive: %v",
					x, y, err)
			}
		}
	}
}
