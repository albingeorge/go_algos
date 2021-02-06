package search

import (
	"errors"
)

//Binary ..binary search
func Binary(input []int, element int) (int, error) {
	pos := search(input, 0, len(input)-1, element)
	if pos != -1 {
		return pos, nil
	}
	return -1, errors.New("not found")
}

func search(input []int, left int, right int, element int) int {
	if right >= left {
		mid := (left + right) / 2

		if element == input[mid] {
			return mid
		}

		if element > input[mid] {
			return search(input, mid+1, right, element)
		}

		return search(input, left, mid-1, element)
	}

	return -1
}
