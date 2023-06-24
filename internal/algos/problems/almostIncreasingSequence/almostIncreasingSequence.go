package almostincreasingsequence

import "fmt"

func AlmostIncreasingSequence(sequence []int) bool {
	return check(sequence, 0)
}

func check(sequence []int, count int) bool {
	fmt.Println(sequence)
	if count > 1 {
		return false
	}

	var prev int
	for i, val := range sequence {
		if i == 0 {
			prev = val
			continue
		}
		if prev >= val {
			fmt.Printf("i: %v\n", i)
			return check(remove(sequence, i), count+1) || check(remove(sequence, i-1), count+1)
		}
		prev = val
	}
	fmt.Printf("Returning true for sequence %v", sequence)
	return true
}

func remove(seq []int, i int) []int {
	out := []int{}
	for k, v := range seq {
		if k != i {
			out = append(out, v)
		}
	}

	return out
}
