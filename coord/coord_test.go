package coord_test

import (
	"fmt"
	"testing"

	"github.com/alcortesm/conway/conway"
	"github.com/alcortesm/conway/coord"
)

func Test(t *testing.T) {
	for _, f := range []struct {
		x, y uint
	}{
		{x: 0, y: 0},
		{x: 0, y: 1},
		{x: 1, y: 0},
		{x: 1, y: 1},
		{x: 300, y: 12},
	} {
		name := fmt.Sprintf("(%d,%d)", f.x, f.y)
		t.Run(name, func(*testing.T) {
			c := coord.New(f.x, f.y)
			conway.TestCoord(t, c, f.x, f.y)
		})
	}
}
