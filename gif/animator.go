package gif

import (
	"image"
	"image/color"
	"image/gif"
	"io"

	"github.com/alcortesm/conway/conway"
)

// Animator represents a collection of grids that will be rendered as a GIF.
// Implements the conway.Animator interface.
type Animator struct {
	howManyGrids       int
	delayBetweenFrames int // 100ths of second
}

// NewAnimator returns a new Animator with the given delay between frames
// in 100ths of seconds.
func NewAnimator(d int) *Animator {
	return &Animator{
		delayBetweenFrames: d,
	}
}

// Add adds a grid to the collection be used as a photogram in the animate method.
// Implements conway.Animator.
func (a *Animator) Add(g conway.Grid) {
	a.howManyGrids++
}

// Encode creates an animation of all the added photograms and store it in
// the given writer.
func (a *Animator) Encode(w io.Writer) error {
	return gif.EncodeAll(w, a.gif())
}

func (a *Animator) gif() *gif.GIF {
	return &gif.GIF{
		Image:     []*image.Paletted{createRedImage(), createWhiteImage()}, // TODO
		Delay:     a.delay(),
		LoopCount: 0,
		Disposal:  nil,
		Config: image.Config{
			ColorModel: nil,
			Width:      100,
			Height:     100,
		},
		BackgroundIndex: 0,
	}
}

func (a *Animator) delay() []int {
	ret := make([]int, a.howManyGrids)
	for i := range ret {
		ret[i] = a.delayBetweenFrames
	}
	return ret
}

func createRedImage() *image.Paletted {
	const width = 100
	const height = 100
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
			G: 0,
			B: 0,
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
			if r%2 == 0 {
				if c%2 == 0 {
					i.Pix[r*width+c] = white
				} else {
					i.Pix[r*width+c] = black
				}
			} else {
				if c%2 == 0 {
					i.Pix[r*width+c] = black
				} else {
					i.Pix[r+width+c] = white
				}
			}
		}
	}
	return i
}

func createWhiteImage() *image.Paletted {
	const width = 100
	const height = 100
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
			if r%2 == 0 {
				if c%2 == 0 {
					i.Pix[r*width+c] = white
				} else {
					i.Pix[r*width+c] = black
				}
			} else {
				if c%2 == 0 {
					i.Pix[r*width+c] = black
				} else {
					i.Pix[r+width+c] = white
				}
			}
		}
	}
	return i
}
