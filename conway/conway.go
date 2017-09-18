package conway

type Universe interface {
	Status() *Grid
	Tick()
}

type Animator interface {
	Add(Grid)
	Animate(file string)
}

// Grid represents an inmutable snapshot of a universe.
type Grid interface {
	// Width returns the width of the universe (number of cells).
	Width() int
	// Height returns the height of the universe (number of cells).
	Height() int
	// IsAlive returns if the cell at coordinates x, y is alive.
	// Returns an error if x or y are out of bounds.
	IsAlive(Coord) (bool, error)
}

// Coord identifies a cell in a 2D coordinate grid.
type Coord interface {
	// X returns the abscissa.
	X() int
	// Y returns the ordinate.
	Y() int
}
