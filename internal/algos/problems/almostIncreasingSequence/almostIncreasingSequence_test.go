package almostincreasingsequence

import "testing"

func TestAlmostIncreasingSequence(t *testing.T) {
	result := AlmostIncreasingSequence([]int{1, 3, 2, 1})

	if result == true {
		t.Errorf("expected false, got true")
	}
}
