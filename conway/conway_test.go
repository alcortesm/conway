package conway_test

import (
	"fmt"
	"reflect"
	"sort"
	"testing"

	"github.com/alcortesm/conway/conway"
	"github.com/alcortesm/conway/coord"
	"github.com/alcortesm/conway/grid"
)

func TestCoordEqualDifferentCoords(t *testing.T) {
	for _, tt := range []struct {
		a, b conway.Coord
	}{
		{coord.New(0, 0), coord.New(1, 1)},
		{coord.New(0, 0), coord.New(0, 1)},
		{coord.New(0, 0), coord.New(1, 0)},
		{coord.New(0, 0), coord.New(1234, 4321)},
		{coord.New(42, 24), coord.New(24, 42)},
	} {
		a := fmt.Sprintf("(%d, %d)", tt.a.X(), tt.a.Y())
		b := fmt.Sprintf("(%d, %d)", tt.b.X(), tt.b.Y())
		if conway.CoordEqual(tt.a, tt.b) {
			t.Errorf("CoordEqual(%s, %s) expected false, got true", a, b)
		}
		if conway.CoordEqual(tt.b, tt.a) {
			t.Errorf("CoordEqual(%s, %s) expected false, got true", b, a)
		}
	}
}

func TestCoordEqualSameCoords(t *testing.T) {
	for _, tt := range []struct {
		a, b conway.Coord
	}{
		{coord.New(0, 0), coord.New(0, 0)},
		{coord.New(1, 0), coord.New(1, 0)},
		{coord.New(0, 1), coord.New(0, 1)},
		{coord.New(1, 1), coord.New(1, 1)},
		{coord.New(42, 24), coord.New(42, 24)},
		{coord.New(5000, 12), coord.New(5000, 12)},
	} {
		a := fmt.Sprintf("(%d, %d)", tt.a.X(), tt.a.Y())
		b := fmt.Sprintf("(%d, %d)", tt.b.X(), tt.b.Y())
		if !conway.CoordEqual(tt.a, tt.b) {
			t.Errorf("CoordEqual((a)%s, (b)%s) expected true, got false", a, b)
		}
		if !conway.CoordEqual(tt.a, tt.b) {
			t.Errorf("CoordEqual((b)%s, (a)%s) expected true, got false", b, a)
		}
	}
}

func TestEvolve(t *testing.T) {
	for _, tt := range []struct {
		name     string
		old      conway.Grid
		expected []conway.Coord
	}{
		{
			name:     "empty",
			old:      checkNewGrid(t, 3, 3, []conway.Coord{}),
			expected: []conway.Coord{},
		}, {
			name: "alives: (0,0)",
			old: checkNewGrid(t, 3, 3, []conway.Coord{
				coord.New(0, 0),
			}),
			expected: []conway.Coord{},
		}, {
			name: "only center is alive",
			old: checkNewGrid(t, 3, 3, []conway.Coord{
				coord.New(1, 1),
			}),
			expected: []conway.Coord{},
		}, {
			name: "alives: (0,0), (2,2)",
			old: checkNewGrid(t, 3, 3, []conway.Coord{
				coord.New(0, 0),
				coord.New(2, 2),
			}),
			expected: []conway.Coord{},
		}, {
			name: "alives: (0,0), (0,2), (2,2)",
			old: checkNewGrid(t, 3, 3, []conway.Coord{
				coord.New(0, 0),
				coord.New(0, 2),
				coord.New(2, 2),
			}),
			expected: []conway.Coord{
				coord.New(1, 1),
			},
		}, {
			name: "alives: (0,0), (0,1), (1,0)",
			old: checkNewGrid(t, 3, 3, []conway.Coord{
				coord.New(0, 0),
				coord.New(0, 1),
				coord.New(1, 0),
			}),
			expected: []conway.Coord{
				coord.New(0, 0),
				coord.New(0, 1),
				coord.New(1, 0),
				coord.New(1, 1),
			},
		}, {
			name: "alives: (0,0), (0,1), (0,2), (1,0), (1, 2)",
			old: checkNewGrid(t, 3, 3, []conway.Coord{
				coord.New(0, 0),
				coord.New(0, 1),
				coord.New(0, 2),
				coord.New(1, 0),
				coord.New(1, 2),
			}),
			expected: []conway.Coord{
				coord.New(0, 0),
				coord.New(0, 2),
				coord.New(1, 0),
				coord.New(1, 2),
			},
		}, {
			name: "alives: (0,0), (0,1), (0,2), (1,0), (2, 1), (2,2)",
			old: checkNewGrid(t, 3, 3, []conway.Coord{
				coord.New(0, 0),
				coord.New(0, 1),
				coord.New(0, 2),
				coord.New(1, 0),
				coord.New(2, 1),
				coord.New(2, 2),
			}),
			expected: []conway.Coord{
				coord.New(0, 0),
				coord.New(0, 1),
				coord.New(1, 0),
				coord.New(2, 1),
				coord.New(2, 2),
			},
		}, {
			name: "all alive",
			old: checkNewGrid(t, 3, 3, []conway.Coord{
				coord.New(0, 0),
				coord.New(0, 1),
				coord.New(0, 2),
				coord.New(1, 0),
				coord.New(1, 1),
				coord.New(1, 2),
				coord.New(2, 0),
				coord.New(2, 1),
				coord.New(2, 2),
			}),
			expected: []conway.Coord{
				coord.New(0, 0),
				coord.New(0, 2),
				coord.New(2, 0),
				coord.New(2, 2),
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			obtained, err := conway.Evolve(tt.old)
			if err != nil {
				t.Fatal(err)
			}
			sort.Sort(byXY{obtained, tt.old.Width()})
			sort.Sort(byXY{tt.expected, tt.old.Width()})
			if !reflect.DeepEqual(obtained, tt.expected) {
				t.Errorf("\nexpected: %#v\nobtained: %#v",
					tt.expected, obtained)
			}
		})
	}
}

func checkNewGrid(t *testing.T, w, h int, alives []conway.Coord) conway.Grid {
	t.Helper()
	g, err := grid.New(w, h, alives)
	if err != nil {
		t.Fatal(err)
	}
	return g
}

type byXY struct {
	alives []conway.Coord
	w      int
}

func (a byXY) Len() int { return len(a.alives) }

func (a byXY) Swap(i, j int) {
	a.alives[i], a.alives[j] = a.alives[j], a.alives[i]
}

func (a byXY) Less(i, j int) bool {
	posI := a.alives[i].X()*a.w + a.alives[i].Y()
	posJ := a.alives[j].X()*a.w + a.alives[j].Y()
	return posI < posJ
}
