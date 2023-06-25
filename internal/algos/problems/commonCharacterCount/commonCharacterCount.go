package commoncharactercount

import (
	"fmt"
	"sort"
)

func Commoncharactercount(s1 string, s2 string) int {
	count := 0
	s1rune := []rune(s1)
	sort.Slice(s1rune, func(i int, j int) bool { return s1rune[i] < s1rune[j] })

	s2rune := []rune(s2)
	sort.Slice(s2rune, func(i int, j int) bool { return s2rune[i] < s2rune[j] })

	// Ensure s1 starts with the smaller alplabet
	if s1rune[0] > s2rune[0] {
		temp := s1rune
		s1rune = s2rune
		s2rune = temp
	}

	i, j := 0, 0
	for {
		if i >= len(s1rune) || j >= len(s2rune) {
			break
		}
		if s1rune[i] == s2rune[j] {
			fmt.Printf("i: %v\n", i)
			end1 := findEnpoint(s1rune, s1rune[i], i)
			end2 := findEnpoint(s2rune, s2rune[j], j)
			diff1 := end1 - i
			diff2 := end2 - j
			if diff1 < diff2 {
				count = count + diff1
			} else {
				count = count + diff2
			}
			i = end1
			j = end2
		} else {
			if s1rune[i] < s2rune[j] {
				i = findEnpoint(s1rune, s1rune[i], i)
			} else {
				j = findEnpoint(s2rune, s2rune[j], j)
			}
		}
		fmt.Printf("Next i: %v; Next j: %v\n", i, j)
	}

	return count
}

func findEnpoint(input []rune, element rune, start int) int {
	i := start
	for ; i < len(input); i++ {
		if input[i] != element {
			break
		}
	}
	return i
}
