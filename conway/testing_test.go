package conway_test

import (
	"testing"

	"github.com/alcortesm/conway/conway"
	"github.com/alcortesm/conway/coord"
)

func TestTestCoordOK(t *testing.T) {
	tt := new(testing.T) // undocumented ctor of testing.T
	conway.TestCoord(tt, coord.New(4, 7), 4, 7)
	if tt.Failed() {
		t.Error()
	}
}

func TestTestCoordFail(t *testing.T) {
	for _, tt := range []struct {
		name string
		c    conway.Coord
		x, y uint
	}{
		{"wrong x", coord.New(4, 7), 0, 7},
		{"wrong y", coord.New(4, 7), 4, 0},
		{"both wrong", coord.New(4, 7), 0, 0},
	} {
		t.Run(tt.name, func(t *testing.T) {
			mockedT := new(testing.T) // undocumented ctor of testing.T
			conway.TestCoord(mockedT, tt.c, tt.x, tt.y)
			if !mockedT.Failed() {
				t.Error()
			}
		})
	}
}
