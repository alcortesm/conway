package conway_test

import (
	"testing"

	"github.com/alcortesm/conway/conway"
	"github.com/alcortesm/conway/coord"
)

func TestTestCoordOK(t *testing.T) {
	mockT := new(testing.T) // undocumented ctor of testing.T
	conway.TestCoord(mockT, coord.New(4, 7), 4, 7)
	if mockT.Failed() {
		t.Error()
	}
}

func TestTestCoordFail(t *testing.T) {
	for _, tt := range []struct {
		name string
		c    conway.Coord
		x, y int
	}{
		{"wrong x", coord.New(4, 7), 0, 7},
		{"wrong y", coord.New(4, 7), 4, 0},
		{"both wrong", coord.New(4, 7), 0, 0},
	} {
		t.Run(tt.name, func(t *testing.T) {
			mockT := new(testing.T) // undocumented ctor of testing.T
			conway.TestCoord(mockT, tt.c, tt.x, tt.y)
			if !mockT.Failed() {
				t.Error()
			}
		})
	}
}
