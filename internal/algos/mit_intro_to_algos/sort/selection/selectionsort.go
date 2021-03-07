package selection

/*
Selection ... Sorts an array using selection sort algorithm

Time Complexity: O(n2)
Space Complexity: O(1)
*/
func Selection(input []int) ([]int, error) {
	n := len(input)
	for i := 0; i < n; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if input[j] < input[minIndex] {
				minIndex = j
			}
		}
		if i != minIndex {
			temp := input[i]
			input[i] = input[minIndex]
			input[minIndex] = temp
		}
	}
	return input, nil
}
