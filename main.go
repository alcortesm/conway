package main

const (
	ticks = 15
	width = 10
	height = 10
)

type Universe interface {
	Status() Grid
	Tick()
}

type Animator interface {
	Add(Grid)
	Animate(file string)
}

type Grid interface {
	Width() int
	Height() int
	Get(x, y int) (bool, error)
	Set(x, y int, v bool) error
}

func main() {
	u := ConwayUniverse{}
	a := GifAnimator{}

	for i:=0; i<ticks; i++ {
		s := u.pic()
		a.add(s)
		u.tick()
	}

	a.animate("conway.gif")
}

