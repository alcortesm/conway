package main

import (
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

	a.Animate("conway.gif")
}
