package badstrings

import "testing"

func TestBadstringsForGoodString(t *testing.T) {
	input := "aeioup??"
	result := Badstrings(input)

	if result != 1 {
		t.Errorf("Bad string failed, expected: %v; got: %v", 1, result)
	}
}

func TestBadstringsForBadString(t *testing.T) {
	input := "bcdaeiou??"
	result := Badstrings(input)

	if result != 0 {
		t.Errorf("Bad string failed, expected: %v; got: %v", 1, result)
	}
}
