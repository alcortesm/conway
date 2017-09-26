package conway_test

import (
	"fmt"
	"testing"

	"github.com/alcortesm/conway/conway"
	"github.com/alcortesm/conway/coord"
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
