package conway

import (
	"testing"
)

// TestCoord is a test helper that checks if the coordinate c has x and y
// as its components.
func TestCoord(t *testing.T, c Coord, x, y uint) {
	t.Helper()
	if x != c.X() {
		t.Errorf("wrong x value: expected %d, got %d", x, c.X())
	}
	if y != c.Y() {
		t.Errorf("wrong y value: expected %d, got %d", y, c.Y())
	}
}
