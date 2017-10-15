package grid_test

import (
	"fmt"
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

const alivePercent uint = 20

func BenchmarkSizeSmall(b *testing.B) {
	var side uint = 10
	alives, err := someAlives(side, alivePercent)
	if err != nil {
		b.Fatalf("cannot create random alives: %v", err)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		benchmarkSize(b, side, alives)
	}
}

func someAlives(side, percent uint) ([]conway.Coord, error) {
	if percent > 100 {
		return nil, fmt.Errorf("percent must be <= 100, was %d", percent)
	}
	n := side * side * percent / 100
	alives := make([]conway.Coord, 0, n)
	var x, y, cont uint
loops:
	for x = 0; x < side; x++ {
		for y = 0; y < side; y++ {
			if cont == n {
				break loops
			}
			alives = append(alives, coord.New(x, y))
			cont++
		}
	}
	return alives, nil
}

func benchmarkSize(b *testing.B, side uint, alives []conway.Coord) {
	g, err := grid.New(side, side, alives)
	if err != nil {
		b.Fatalf("cannot create grid: %v", err)
	}
	var x, y uint
	for x = 0; x < side; x++ {
		for y = 0; y < side; y++ {
			ia, err := g.IsAlive(coord.New(x, y))
			if err != nil {
				b.Fatalf("checking if (%d, %d) is alive: %v",
					x, y, err)
			}
			_ = ia
		}
	}
}

func BenchmarkSizeMedium(b *testing.B) {
	var side uint = 100
	alives, err := someAlives(side, alivePercent)
	if err != nil {
		b.Fatalf("cannot create random alives: %v", err)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		benchmarkSize(b, side, alives)
	}
}

func BenchmarkSizeBig(b *testing.B) {
	var side uint = 1000
	alives, err := someAlives(side, alivePercent)
	if err != nil {
		b.Fatalf("cannot create random alives: %v", err)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		benchmarkSize(b, side, alives)
	}
}
