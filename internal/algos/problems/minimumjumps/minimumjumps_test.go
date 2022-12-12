package minimumjumps

import "testing"

func TestMinimumJumps(t *testing.T) {

	type test struct {
		input    []int
		expected int
	}

	tests := []test{
		{
			input:    []int{1, 3, 5, 8, 9, 2, 6, 7, 6, 8, 9},
			expected: 3,
		},
		{
			input:    []int{1, 4, 3, 2, 6, 7},
			expected: 2,
		},
	}

	for _, tc := range tests {
		result := MinimumJumps(tc.input)
		if result != tc.expected {
			t.Errorf("Minimum jumps failed, input: %v; expected: %v; got: %v", tc.input, tc.expected, result)
		}
	}

}
