package conway

import (
	"fmt"
	"testing"
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
