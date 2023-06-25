package commoncharactercount

import "testing"

func TestCommoncharactercount(t *testing.T) {
	expected := 3
	input1, input2 := "aabcc", "adcaa"
	got := Commoncharactercount(input1, input2)

	if expected != got {
		t.Errorf("Expected: %v; got: %v", expected, got)
	}
}
