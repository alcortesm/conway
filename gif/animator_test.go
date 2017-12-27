package gif

import (
	"io/ioutil"
	"math/rand"
	"testing"

	"github.com/alcortesm/conway/conway"
	"github.com/alcortesm/conway/coord"
	"github.com/alcortesm/conway/grid"
)

func Test(t *testing.T) {
	f, err := ioutil.TempFile("", "test_animator_")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	t.Logf("file written at %s", f.Name())

	delay := 10
	resolution := 10
	a, err := NewAnimator(delay, resolution)
	if err != nil {
		t.Fatal(err)
	}

	ticks := 10
	var width uint = 50
	var height uint = 50
	for i := 0; i < ticks; i++ {
		g, err := grid.New(width, height, random(width, height))
		if err != nil {
			t.Fatal(err)
		}
		a.Add(g)
	}
	if err := a.Encode(f); err != nil {
		t.Error(err)
	}
}

func random(w, h uint) []conway.Coord {
	count := 500
	ret := []conway.Coord{}
	for i := 0; i < count; i++ {
		x := uint(rand.Intn(int(w)))
		y := uint(rand.Intn(int(h)))
		ret = append(ret, coord.New(x, y))
	}
	return ret
}
