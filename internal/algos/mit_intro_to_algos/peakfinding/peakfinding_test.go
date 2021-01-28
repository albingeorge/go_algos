package peakfinding

import "testing"

func TestPeakFinding(t *testing.T) {
	input := []int{1, 3, 5, 5, 4, 6, 3, 2, 1, 4, 7, 3, 4, 2, 5}
	pos := peakFindRecursive(input, 0, len(input)-1)

	if pos != 3 {
		t.Errorf("Peak finding failed, expected: %v; got: %v", 3, pos)
	}
}
