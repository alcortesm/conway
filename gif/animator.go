package gif

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"

	"github.com/alcortesm/conway/conway"
	"github.com/alcortesm/conway/coord"
)

// Animator represents a collection of grids that will be rendered as a GIF.
// Implements the conway.Animator interface.  The origin of coordinates is
// set at the upper left corner and cells with negatives coordinates are
// ignored.
// The zero value of this type is not safe, use the function NewAnimator below.
type Animator struct {
	delayBetweenFrames int // 100ths of second
	addedGrids         []conway.Grid
	resolution         int // pixels per cell
	// TODO add the colors to use in the animation
}

// NewAnimator returns a new Animator with the given delay d between frames
// in 100ths of seconds and the given resolution r in pixels per cell.
func NewAnimator(d int, r int) (*Animator, error) {
	if d < 1 {
		return nil, fmt.Errorf("delay has to be > 0, got %d", d)
	}
	if r < 1 {
		return nil, fmt.Errorf("resolution has to be > 0, got %d", r)
	}
	return &Animator{
		delayBetweenFrames: d,
		resolution:         r,
	}, nil
}

// Add adds a grid to the collection to be used as a photogram in the animate method.
// Implements conway.Animator.
func (a *Animator) Add(g conway.Grid) error {
	if len(a.addedGrids) > 0 {
		first := a.addedGrids[0]
		if g.Width() != first.Width() || g.Height() != first.Height() {
			return fmt.Errorf(
				"unexpected size, expected (%d, %d), got (%d, %d)",
				first.Width(), first.Height(), g.Width(), g.Height())
		}
	}
	a.addedGrids = append(a.addedGrids, g)
	return nil
}

// Encode creates an animation of all the added photograms and store it in
// the given writer.
func (a *Animator) Encode(w io.Writer) error {
	if len(a.addedGrids) == 0 {
		return fmt.Errorf("Add some grids first")
	}
	return gif.EncodeAll(w, a.gif())
}

func (a *Animator) gif() *gif.GIF {
	first := a.addedGrids[0]
	widthPixels := first.Width() * a.resolution
	heightPixels := first.Height() * a.resolution
	return &gif.GIF{
		Image:     a.images(),
		Delay:     a.delay(),
		LoopCount: 0,
		Disposal:  nil,
		Config: image.Config{
			ColorModel: nil,
			Width:      widthPixels,
			Height:     heightPixels,
		},
		BackgroundIndex: 0,
	}
}

// Images returs an array with the added grids turned into images.
func (a *Animator) images() []*image.Paletted {
	ret := make([]*image.Paletted, len(a.addedGrids))
	for i := range ret {
		ret[i] = a.gridToImage(a.addedGrids[i])
	}
	return ret
}

// GridToImage transforms a grid into an image.
func (a *Animator) gridToImage(g conway.Grid) *image.Paletted {
	width := g.Width() * a.resolution
	height := g.Height() * a.resolution
	r := image.Rectangle{
		Min: image.Point{
			X: 0,
			Y: 0,
		},
		Max: image.Point{
			X: width,
			Y: height,
		},
	}
	p := color.Palette{
		color.RGBA{
			R: 255,
			G: 255,
			B: 255,
			A: 255,
		},
		color.RGBA{
			R: 0,
			G: 0,
			B: 0,
			A: 255,
		},
	}
	i := image.NewPaletted(r, p)
	const white = 0
	const black = 1
	for c := 0; c < width; c++ {
		for r := 0; r < height; r++ {
			cell := coord.New(c/a.resolution, r/a.resolution)
			isAlive, err := g.IsAlive(cell)
			if err != nil {
				panic(err) // unreachable
			}
			if isAlive {
				i.Pix[r*width+c] = black
			} else {
				i.Pix[r*width+c] = white
			}
		}
	}
	return i
}

func (a *Animator) delay() []int {
	ret := make([]int, len(a.addedGrids))
	for i := range ret {
		ret[i] = a.delayBetweenFrames
	}
	return ret
}
