package search

import (
	"errors"
	"math"
)

func Jump(input []int, element int) (int, error) {
	size := len(input)

	// Step size - root of the length of array
	root := int(math.Sqrt(float64(size)))
	i := 0

	for i < size && input[i] < element {
		i += root
	}

	if i <= size {
		// Get the previous window
		j := i - root
		// If the previos window starts from a negative index set it to start
		if j < 0 {
			j = 0
		}
		for ; j < size; j++ {
			if input[j] == element {
				return j, nil
			}
		}
	}

	return -1, errors.New("not found")
}
