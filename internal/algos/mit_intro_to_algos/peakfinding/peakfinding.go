package peakfinding

import "errors"

//Peakfind ...find the peak in a slice of ints
func Peakfind(input []int) int {
	// input := []int{1, 3, 5, 5, 4, 6, 3, 2, 1, 4, 7, 3, 4, 2, 5}

	last := len(input) - 1
	pos := peakFindRecursive(input, 0, last)

	return pos
}

//PeakFindBruteForce ... peak find by brute force
func PeakFindBruteForce(input []int) (int, error) {
	if len(input) == 0 {
		return 0, errors.New("empty input")
	} else if len(input) == 1 {
		return 0, nil
	} else if len(input) == 2 {
		if input[0] < input[1] {
			return 1, nil
		}
		return 0, nil
	}

	for i, v := range input {
		// If left edge is a peak
		if i == 0 {
			if v > input[i+1] {
				return i, nil
			}
			continue
		}
		// If right edge is reached
		if i == len(input)-1 {
			continue
		}

		if input[i-1] <= v && v >= input[i+1] {
			return i, nil
		}
	}

	return len(input) - 1, nil
}

func peakFindRecursive(input []int, start int, end int) int {
	mid := start + ((end - start) / 2)

	if (end - start) == 0 {
		// If there's only one element
		return start
	} else if end-start == 1 {
		// If there are only 2 elements
		if input[start] < input[end] {
			return end
		}
		return start
	}

	if input[mid-1] <= input[mid] && input[mid] >= input[mid+1] {
		return mid
	} else if input[mid-1] < input[mid] {
		return peakFindRecursive(input, mid, end)
	} else {
		return peakFindRecursive(input, start, mid-1)
	}
}
