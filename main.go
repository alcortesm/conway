package main

type Universe interface {
	Tick()
	Snapshot() [][]bool
}

type Animator interface {
	AddFrame([][]bool)
	Animate()
}

const ticks = 10

func main() {
	var u Universe
	var a Animator
	for i:=0; i<ticks; i++ {
		current := u.Snapshot()
		a.AddFrame(current)
		u.Tick()
	}
	a.Animate()
}
