package peakfinding

//Peakfind ...find the peak in a slice of ints
func Peakfind(input []int) int {
	// input := []int{1, 3, 5, 5, 4, 6, 3, 2, 1, 4, 7, 3, 4, 2, 5}

	last := len(input) - 1
	pos := peakFindRecursive(input, 0, last)

	return pos
}

func peakFindRecursive(input []int, start int, end int) int {
	mid := start + ((end - start) / 2)

	if (end - start) == 0 {
		return start
	} else if end-start == 1 {
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
