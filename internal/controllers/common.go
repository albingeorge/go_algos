package controllers

import (
	"strconv"
	"strings"
)

func formatInput(inputString string) []int {
	splitString := strings.Split(inputString, ",")
	var input []int

	for _, s := range splitString {
		inputString := strings.Trim(s, " ")
		converted, _ := strconv.Atoi(inputString)
		input = append(input, converted)
	}
	return input
}
