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
	testSizeOK(t, grid.MinWidth, grid.MinHeight)
	testSizeOK(t, grid.MinWidth, grid.MinHeight+1)
	testSizeOK(t, grid.MinWidth+1, grid.MinHeight)
	testSizeOK(t, grid.MinWidth+1, grid.MinHeight+1)
	testSizeOK(t, grid.MinWidth+100, grid.MinHeight+200)
	testSizeOK(t, grid.MinWidth+200, grid.MinHeight+100)
	testSizeOK(t, grid.MinWidth+200, grid.MinHeight+200)
	testSizeOK(t, grid.MinWidth+20000, grid.MinHeight+20000)
}

func testSizeOK(t *testing.T, width, height int) {
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
	testSizeError(t, grid.MinWidth-1, grid.MinHeight)
	testSizeError(t, grid.MinWidth, grid.MinHeight-1)
	testSizeError(t, grid.MinWidth-1, grid.MinHeight-1)
	testSizeError(t, 0, 0)
	testSizeError(t, -1, grid.MinHeight)
	testSizeError(t, grid.MinWidth, -1)
	testSizeError(t, -1, -1)
}

func testSizeError(t *testing.T, width, height int) {
	t.Helper()
	_, err := grid.New(width, height, nil)
	if err == nil {
		t.Errorf("new grid was supposed to fail and it did not")
	}
}

func TestNewErrorWithAliveOutOfBounds(t *testing.T) {
	for _, tt := range []struct {
		name   string
		w, h   int
		alives []conway.Coord // some of them will be out of bounds
	}{
		{"x out of bounds", 3, 3, []conway.Coord{
			coord.New(3, 1), // out of bounds
		}},
		{"y out of bounds", 3, 3, []conway.Coord{
			coord.New(1, 3), // out of bounds
		}},
		{"both out of bounds", 3, 3, []conway.Coord{
			coord.New(3, 3), // out of bounds
		}},
		{"all out of bounds", 3, 3, []conway.Coord{
			coord.New(1, 3), // out of bounds
			coord.New(3, 1), // out of bounds
			coord.New(3, 3), // out of bounds
		}},
		{"some out of bounds", 3, 3, []conway.Coord{
			coord.New(0, 0),
			coord.New(3, 1), // out of bounds
			coord.New(1, 2),
			coord.New(14, 2), // out of bounds
		}},
		{"negative x", 3, 3, []conway.Coord{
			coord.New(-1, 1), // out of bounds
		}},
		{"negative y", 3, 3, []conway.Coord{
			coord.New(1, -1), // out of bounds
		}},
		{"both negative", 3, 3, []conway.Coord{
			coord.New(-1, -1), // out of bounds
		}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			_, err := grid.New(tt.w, tt.h, tt.alives)
			if err == nil {
				t.Errorf("new grid was supposed to fail but didn't")
			}
		})
	}
}

func TestIsAlive(t *testing.T) {
	for _, tt := range []struct {
		name   string
		w, h   int
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

func testIsAlive(t *testing.T, w, h int, alives []conway.Coord) {
	g, err := grid.New(w, h, alives)
	if err != nil {
		t.Fatalf("cannot create grid: %v", err)
	}
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
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

func TestIsAliveOutOfBounds(t *testing.T) {
	g, err := grid.New(4, 4, []conway.Coord{
		coord.New(1, 0),
		coord.New(2, 1),
		coord.New(0, 2),
	})
	if err != nil {
		t.Fatalf("cannot create grid: %v", err)
	}
	if _, err = g.IsAlive(coord.New(7, 7)); err == nil {
		t.Error()
	}
}

func BenchmarkSize(b *testing.B) {
	for _, side := range []int{300, 600, 900, 1200, 1500, 1800, 2100, 2400, 2700} {
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

func allCoords(side int) []conway.Coord {
	all := make([]conway.Coord, 0, side*side)
	for x := 0; x < side; x++ {
		for y := 0; y < side; y++ {
			all = append(all, coord.New(x, y))
		}
	}
	return all
}

func percentOf(p, total int) int {
	return total * p / 100
}

func randomSubSlice(n int, all []conway.Coord) []conway.Coord {
	ret := make([]conway.Coord, n)
	randomIndexes := rand.Perm(len(all))
	for i, v := range randomIndexes[:n] {
		ret[i] = all[v]
	}
	return ret
}

func createGridAndCheckAllCells(b *testing.B, side int, alives []conway.Coord) {
	g, err := grid.New(side, side, alives)
	if err != nil {
		b.Fatalf("cannot create grid: %v", err)
	}
	for x := 0; x < side; x++ {
		for y := 0; y < side; y++ {
			if _, err := g.IsAlive(coord.New(x, y)); err != nil {
				b.Fatalf("checking if (%d, %d) is alive: %v",
					x, y, err)
			}
		}
	}
}
