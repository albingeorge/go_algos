package bubble

import "testing"

func TestBubbleSort(t *testing.T) {
	input := []int{4, 2, 1, 3, 5, 6}
	expected := []int{1, 2, 3, 4, 5, 6}
	result, _ := Bubble(input)

	comp := true
	if len(input) != len(expected) {
		comp = false
	}
	if comp == true {
		for i, v := range expected {
			if v != result[i] {
				comp = false
			}
		}
	}

	if comp != true {
		t.Errorf("Bubble sorting failed, expected: %v; got: %v", expected, result)
	}
}
