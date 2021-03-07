package bubble

/*
Bubble ... Implementes bubble sort
*/
func Bubble(input []int) ([]int, error) {
	n := len(input)
	var temp int
	for i := 0; i < n; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if input[j] > input[j+1] {
				temp = input[j]
				input[j] = input[j+1]
				input[j+1] = temp
				swapped = true
			}
		}
		if swapped == false {
			break
		}
	}
	return input, nil
}
