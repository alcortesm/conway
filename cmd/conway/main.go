package main

import (
	"log"
	"os"

	"github.com/alcortesm/conway/conway"
)

const (
	ticks  = 15
	width  = 10
	height = 10
)

func main() {
	var u conway.Universe
	var a conway.Animator

	for i := 0; i < ticks; i++ {
		grid := u.Status()
		a.Add(grid)
		u.Tick()
	}

	f, err := os.Create("conway.gif")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := a.Encode(f); err != nil {
		log.Fatal(err)
	}
}
