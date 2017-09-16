package conway

type Universe interface {
	Status() *Grid
	Tick()
}

type Animator interface {
	Add(Grid)
	Animate(file string)
}

// Grid represents a snapshot of a universe.
type Grid interface {
	// Width returns the width of the universe (number of cells).
	Width() int
	// Height returns the height of the universe (number of cells).
	Height() int
	// Get returns if the cell at coordinates x, y is alive.  Returns
	// an error if x or y are out of bounds.
	Get(x, y int) (bool, error)
	// Set sets the cell at coordinates x, y to value v (true for alive,
	// false for dead).  Returns an error if x or y are out of bounds.
	Set(x, y int, v bool) error
}
