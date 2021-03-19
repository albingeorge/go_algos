package search

import "testing"

func TestJumpSearch(t *testing.T) {
	in := []int{2, 3, 4, 6, 7, 8, 9, 10, 14}
	pos, _ := Jump(in, 7)
	if pos != 4 {
		t.Errorf("Jump search failed, expected: %v; got: %v", 1, pos)
	}

	in = []int{1, 2, 3}
	_, err := Jump(in, 4)
	if err == nil {
		t.Error("Jump search failed, error not thrown")
	}

	if err.Error() != "not found" {
		t.Errorf("Jump search failed, invalid error message; expected: %v, received: %v", "not found", err.Error())
	}

}
