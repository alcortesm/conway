package grid

import "fmt"

type Grid struct {
	width, height int
}

const (
	minWidth  = 3
	minHeight = 3
)

func New(width, height int) (*Grid, error) {
	if width < minWidth {
		return nil, fmt.Errorf("width must be >= than %d, was %d",
			minWidth, width)
	}
	if height < minHeight {
		return nil, fmt.Errorf("height must be >= than %d, was %d",
			minHeight, height)
	}
	return &Grid{
		width:  width,
		height: height,
	}, nil
}

func (g *Grid) Width() int {
	return g.width
}

func (g *Grid) Height() int {
	return g.height
}

func (g *Grid) Get(x, y int) (bool, error) {
	return false, fmt.Errorf("TODO")
}

func (g *Grid) Set(x, y int, v bool) error {
	return fmt.Errorf("TODO")
}
