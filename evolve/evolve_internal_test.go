package evolve

import (
	"fmt"
	"testing"

	"github.com/alcortesm/conway/conway"
	"github.com/alcortesm/conway/coord"
	"github.com/alcortesm/conway/grid"
)

func TestLawsCurrentIsAlive(t *testing.T) {
	for _, tt := range []struct {
		neighbours int
		want       bool
	}{
		{neighbours: 0, want: false},
		{neighbours: 1, want: false},
		{neighbours: 2, want: true},
		{neighbours: 3, want: true},
		{neighbours: 4, want: false},
		{neighbours: 5, want: false},
		{neighbours: 6, want: false},
		{neighbours: 7, want: false},
		{neighbours: 8, want: false},
	} {
		name := fmt.Sprintf("%d alive neighbours", tt.neighbours)
		t.Run(name, func(t *testing.T) {
			got := laws(true, tt.neighbours)
			if got != tt.want {
				t.Error()
			}
		})
	}
}

func TestLawsCurrentIsDead(t *testing.T) {
	for _, tt := range []struct {
		neighbours int
		want       bool
	}{
		{neighbours: 0, want: false},
		{neighbours: 1, want: false},
		{neighbours: 2, want: false},
		{neighbours: 3, want: true},
		{neighbours: 4, want: false},
		{neighbours: 5, want: false},
		{neighbours: 6, want: false},
		{neighbours: 7, want: false},
		{neighbours: 8, want: false},
	} {
		name := fmt.Sprintf("%d alive neighbours", tt.neighbours)
		t.Run(name, func(t *testing.T) {
			got := laws(false, tt.neighbours)
			if got != tt.want {
				t.Error()
			}
		})
	}
}

func TestCountNeighbours(t *testing.T) {
	for _, tt := range []struct {
		name    string
		grid    conway.Grid
		current conway.Coord
		want    int
	}{
		{
			name:    "in the middle, empty",
			grid:    empty(t),
			current: coord.New(1, 1),
			want:    0,
		}, {
			name:    "in the middle, full",
			grid:    full(t),
			current: coord.New(1, 1),
			want:    8,
		}, {
			name:    "top left corner, empty",
			grid:    empty(t),
			current: coord.New(0, 0),
			want:    0,
		}, {
			name:    "top left corner, full",
			grid:    full(t),
			current: coord.New(0, 0),
			want:    3,
		}, {
			name:    "top right corner, empty",
			grid:    empty(t),
			current: coord.New(2, 0),
			want:    0,
		}, {
			name:    "top right corner, full",
			grid:    full(t),
			current: coord.New(2, 0),
			want:    3,
		}, {
			name:    "lower left corner, empty",
			grid:    empty(t),
			current: coord.New(0, 2),
			want:    0,
		}, {
			name:    "lower left corner, full",
			grid:    full(t),
			current: coord.New(0, 2),
			want:    3,
		}, {
			name:    "lower right corner, empty",
			grid:    empty(t),
			current: coord.New(2, 2),
			want:    0,
		}, {
			name:    "lower right corner, full",
			grid:    full(t),
			current: coord.New(2, 2),
			want:    3,
		}, {
			name:    "top side, empty",
			grid:    empty(t),
			current: coord.New(1, 0),
			want:    0,
		}, {
			name:    "top side, full",
			grid:    full(t),
			current: coord.New(1, 0),
			want:    5,
		}, {
			name:    "right side, empty",
			grid:    empty(t),
			current: coord.New(2, 1),
			want:    0,
		}, {
			name:    "right side, full",
			grid:    full(t),
			current: coord.New(2, 1),
			want:    5,
		}, {
			name:    "lower side, empty",
			grid:    empty(t),
			current: coord.New(1, 2),
			want:    0,
		}, {
			name:    "lower side, full",
			grid:    full(t),
			current: coord.New(1, 2),
			want:    5,
		}, {
			name:    "left side, empty",
			grid:    empty(t),
			current: coord.New(0, 1),
			want:    0,
		}, {
			name:    "left side, full",
			grid:    full(t),
			current: coord.New(0, 1),
			want:    5,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := countNeighbours(tt.grid, tt.current)
			if got != tt.want {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}

func empty(t *testing.T) conway.Grid {
	t.Helper()
	g, err := grid.New(3, 3, []conway.Coord{})
	if err != nil {
		t.Error(err)
	}
	return g
}

func full(t *testing.T) conway.Grid {
	t.Helper()
	alives := make([]conway.Coord, 0, 9)
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			c := coord.New(x, y)
			alives = append(alives, c)
		}
	}
	g, err := grid.New(3, 3, alives)
	if err != nil {
		t.Error(err)
	}
	return g
}
