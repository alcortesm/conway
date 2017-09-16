package conway

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
