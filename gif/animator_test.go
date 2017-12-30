package gif

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"reflect"
	"testing"

	"github.com/alcortesm/conway/conway"
	"github.com/alcortesm/conway/coord"
	"github.com/alcortesm/conway/grid"
)

func TestNewFailsWithNonPositiveDelay(t *testing.T) {
	for _, delay := range []int{0, -10} {
		name := fmt.Sprintf("delay=%d", delay)
		t.Run(name, func(t *testing.T) {
			resolution := 100
			if _, err := NewAnimator(delay, resolution); err == nil {
				t.Error()
			}
		})
	}
}

func TestNewFailsWithNonPositiveResolutions(t *testing.T) {
	for _, resolution := range []int{0, -10} {
		name := fmt.Sprintf("resolution=%d", resolution)
		t.Run(name, func(t *testing.T) {
			delay := 100
			if _, err := NewAnimator(delay, resolution); err == nil {
				t.Error()
			}
		})
	}
}

func TestEncodeFailsIfNoAdds(t *testing.T) {
	delay, resolution := 10, 100
	a, err := NewAnimator(delay, resolution)
	if err != nil {
		t.Fatal(err)
	}

	var w bytes.Buffer
	if err := a.Encode(&w); err == nil {
		t.Error()
	}
}

func TestAddFailsIfGridsHaveDifferentSizes(t *testing.T) {
	delay, resolution := 10, 100
	a, err := NewAnimator(delay, resolution)
	if err != nil {
		t.Fatal(err)
	}

	width, height, alives := uint(5), uint(7), []conway.Coord{}
	g, err := grid.New(width, height, alives)
	if err != nil {
		t.Fatal(err)
	}
	if err := a.Add(g); err != nil {
		t.Fatal(err)
	}

	width++
	differentSizeGrid, err := grid.New(width, height, alives)
	if err != nil {
		t.Fatal(err)
	}
	if err := a.Add(differentSizeGrid); err == nil {
		t.Error(err)
	}
}

func TestEncode(t *testing.T) {
	for _, tt := range []struct {
		name              string
		delay, resolution int
		grids             []conway.Grid
	}{
		{
			name:  "one empty grid, smallest, resolution 1",
			delay: 10, resolution: 1,
			grids: []conway.Grid{newEmptyGrid(t, 3, 3)},
		}, {
			name:  "one empty grid",
			delay: 10, resolution: 100,
			grids: []conway.Grid{newEmptyGrid(t, 5, 7)},
		}, {
			name:  "two empty grids",
			delay: 10, resolution: 100,
			grids: []conway.Grid{
				newEmptyGrid(t, 5, 7),
				newEmptyGrid(t, 5, 7),
			},
		}, {
			name:  "one full grid",
			delay: 10, resolution: 100,
			grids: []conway.Grid{newFullGrid(t, 5, 7)},
		}, {
			name:  "two full grids",
			delay: 10, resolution: 100,
			grids: []conway.Grid{
				newFullGrid(t, 5, 7),
				newFullGrid(t, 5, 7),
			},
		}, {
			name:  "one empty, one half, one full",
			delay: 10, resolution: 100,
			grids: []conway.Grid{
				newEmptyGrid(t, 5, 7),
				newHalfGrid(t, 5, 7),
				newFullGrid(t, 5, 7),
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGif(t, tt.delay, tt.resolution, tt.grids)
			checkDelay(t, g, len(tt.grids), tt.delay)
			width := tt.grids[0].Width()
			height := tt.grids[0].Height()
			checkConfig(t, g, width, height, tt.resolution)
			checkDisposal(t, g, len(tt.grids))
			checkImage(t, g, tt.resolution, tt.grids)
		})
	}
}

func newEmptyGrid(t *testing.T, w, h uint) conway.Grid {
	t.Helper()
	g, err := grid.New(w, h, []conway.Coord{})
	if err != nil {
		t.Fatal("cannot create empty grid:", err)
	}
	return g
}

func newFullGrid(t *testing.T, w, h uint) conway.Grid {
	t.Helper()
	alives := []conway.Coord{}
	for c := uint(0); c < w; c++ {
		for r := uint(0); r < h; r++ {
			alives = append(alives, coord.New(c, r))
		}
	}
	g, err := grid.New(w, h, alives)
	if err != nil {
		t.Fatal("cannot create full grid:", err)
	}
	return g
}

func newHalfGrid(t *testing.T, w, h uint) conway.Grid {
	t.Helper()
	alives := []conway.Coord{}
	lastWasAlive := false
	for c := uint(0); c < w; c++ {
		for r := uint(0); r < h; r++ {
			if lastWasAlive {
				lastWasAlive = false
				continue
			}
			alives = append(alives, coord.New(c, r))
			lastWasAlive = true
		}
	}
	g, err := grid.New(w, h, alives)
	if err != nil {
		t.Fatal("cannot create half full grid:", err)
	}
	return g
}

func NewGif(t *testing.T, delay, resolution int, grids []conway.Grid) *gif.GIF {
	t.Helper()

	a, err := NewAnimator(delay, resolution)
	if err != nil {
		t.Fatal("cannot create animator:", err)
	}

	for i, g := range grids {
		if err := a.Add(g); err != nil {
			t.Fatalf("cannot add grid #%d: %v", i, err)
		}
	}

	var buf bytes.Buffer
	if err := a.Encode(&buf); err != nil {
		t.Fatal("cannot encode:", err)
	}

	decoded, err := gif.DecodeAll(&buf)
	if err != nil {
		t.Fatal("cannot decode:", err)
	}

	return decoded
}

func checkDelay(t *testing.T, g *gif.GIF, length, value int) {
	t.Helper()
	if len(g.Delay) != length {
		t.Errorf("wrong length, want %d, got %d", length, len(g.Delay))
	}
	for i, d := range g.Delay {
		if d != value {
			t.Errorf("delay #%d: want %d, got %d", i, d, value)
		}
	}
}

func checkConfig(t *testing.T, g *gif.GIF, w, h uint, r int) {
	t.Helper()
	widthPixels := int(w) * r
	heightPixels := int(h) * r
	if g.Config.Width != widthPixels {
		t.Errorf("wrong width, want %d, got %d", widthPixels, g.Config.Width)
	}
	if g.Config.Height != heightPixels {
		t.Errorf("wrong height, want %d, got %d", heightPixels, g.Config.Height)
	}
}

func checkDisposal(t *testing.T, g *gif.GIF, length int) {
	t.Helper()
	if len(g.Disposal) != length {
		t.Errorf("wrong disposal length, want %d, got %d", length, len(g.Disposal))
	}
	for i, e := range g.Disposal {
		if e != 0 {
			t.Errorf("disposal #%d, expected 0, got %d", i, e)
		}
	}
}

func checkImage(t *testing.T, g *gif.GIF, resolution int, grids []conway.Grid) {
	t.Helper()

	if g.BackgroundIndex != 0 {
		t.Errorf("wrong background index: want 0, got %d", g.BackgroundIndex)
	}
	if len(g.Image) != len(grids) {
		t.Errorf("wrong image length:, want %d, got %d", len(grids), len(g.Image))
	}
	palette := color.Palette{
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

	widthPixels := int(grids[0].Width()) * resolution
	heightPixels := int(grids[0].Height()) * resolution
	for i, frame := range g.Image {
		if frame.Stride != widthPixels {
			t.Errorf("wrong Stride in frame #%d: want %d, got %d", i, widthPixels, frame.Stride)
		}
		if expected := image.Rect(0, 0, widthPixels, heightPixels); frame.Rect != expected {
			t.Errorf("wrong rectangle in frame #%d: want %v, got %v", i, expected, frame.Rect)
		}
		if !reflect.DeepEqual(frame.Palette, palette) {
			t.Errorf("wrong palette in frame #%d\nwant %v\ngot  %v", i, palette, frame.Palette)
		}
		for c := 0; c < widthPixels; c++ {
			for r := 0; r < heightPixels; r++ {
				cell := coord.New(uint(c/resolution), uint(r/resolution))
				alive, err := grids[i].IsAlive(cell)
				if err != nil {
					t.Fatalf("cannot check if (%d, %d) is alive in frame #%d: %v",
						cell.X(), cell.Y(), i, err)
				}
				pos := r*frame.Stride + c
				expectedPaletteIndex := uint8(0)
				if alive {
					expectedPaletteIndex = 1
				}
				if frame.Pix[pos] != expectedPaletteIndex {
					t.Errorf("cell (%d, %d) int frame #%d: wrong color, want %d, got %d",
						cell.X(), cell.Y(), i, expectedPaletteIndex, frame.Pix[pos])
				}
			}
		}
	}
}
