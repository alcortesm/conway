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
// Implements the conway.Animator interface.
// The zero value of this type is not safe, use the function NewAnimator below.
type Animator struct {
	delayBetweenFrames int // 100ths of second
	addedGrids         []conway.Grid
	resolution         int // pixels per cell
	// TODO add the colors to use in the animation
}

// NewAnimator returns a new Animator with the given delay between frames
// in 100ths of seconds.
func NewAnimator(d int, r int) (*Animator, error) {
	if r < 1 {
		return nil, fmt.Errorf("resolution has to be > 1, got %d", r)
	}
	return &Animator{
		delayBetweenFrames: d, // TODO add error checking here
		resolution:         r,
	}, nil
}

// Add adds a grid to the collection to be used as a photogram in the animate method.
// Implements conway.Animator.
func (a *Animator) Add(g conway.Grid) {
	// TODO add errors for incompatible grids.
	a.addedGrids = append(a.addedGrids, g)
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
	return &gif.GIF{
		Image:     a.images(),
		Delay:     a.delay(),
		LoopCount: 0,
		Disposal:  nil,
		Config: image.Config{
			ColorModel: nil,
			Width:      int(a.addedGrids[0].Width()) * a.resolution,
			Height:     int(a.addedGrids[0].Height()) * a.resolution,
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
	width := int(g.Width()) * a.resolution
	height := int(g.Height()) * a.resolution
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
			cell := coord.New(uint(c/a.resolution), uint(r/a.resolution))
			isAlive, err := g.IsAlive(cell)
			if err != nil {
				panic(err)
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
