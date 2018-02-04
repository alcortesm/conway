package main

import (
	"log"
	"os"

	"github.com/alcortesm/conway/animator/gif"

	"github.com/alcortesm/conway/conway"
	"github.com/alcortesm/conway/universe/random"
)

const (
	ticks      = 15
	width      = 10
	height     = 10
	nAlives    = width * height / 2
	delay      = 50 // 100ths of seconds
	resolution = 30 // side size of the cell in pixels
)

func main() {
	var u conway.Universe
	var a conway.Animator
	var err error

	u, err = random.New(width, height, nAlives)
	if err != nil {
		log.Fatal("cannot create random universe: ", err)
	}

	a, err = gif.NewAnimator(delay, resolution)
	if err != nil {
		log.Fatal("creating gif animator: ", err)
	}

	for i := 0; i < ticks; i++ {
		grid := u.Status()
		a.Add(grid)
		u.Tick()
	}

	f, err := os.Create(os.TempDir() + "/conway.gif")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := a.Encode(f); err != nil {
		log.Fatal(err)
	}
}
