package grid_test

import (
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
	g, err := grid.New(3, 3, []conway.Coord{
		coord.New(0, 0),
		coord.New(1, 1),
		coord.New(2, 2),
	})
	if err != nil {
		t.Fatalf("cannot create grid: %v", err)
	}
	ia, err := g.IsAlive(coord.New(0, 0))
	if err != nil {
		t.Fatalf("error calling IsAlive: %v", err)
	}
	if ia == false {
		t.Errorf("coordenate (0,0) should be alive but it is dead")
	}
}
