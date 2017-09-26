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
	Width() uint
	// Height returns the height of the universe (number of cells).
	Height() uint
	// IsAlive returns if the cell at coordinates x, y is alive.
	// Returns an error if x or y are out of bounds.
	IsAlive(Coord) (bool, error)
}

// Coord represents the position of a cell in a grid.
type Coord interface {
	// X returns the cell's abscissa.
	X() uint
	// Y returns the cell's ordinate.
	Y() uint
}

func CoordEqual(a, b Coord) bool {
	return a.X() == b.X() && a.Y() == b.Y()
}
