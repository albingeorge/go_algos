package minimumjumps

import "fmt"

func MinimumJumps(arr []int) int {
	//1, 4, 3, 2, 6, 7
	jumps := 0
	for i := 0; i < len(arr)-1; {
		if arr[i] == 0 {
			return -1
		}

		i = i + arr[i]
		fmt.Printf("i: %v\n", i)
		jumps++
	}

	return jumps
}
