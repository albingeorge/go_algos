package controllers

import (
	"strconv"
	"strings"

	"github.com/albingeorge/go_algos/internal/algos/mit_intro_to_algos/peakfinding"
	"github.com/gin-gonic/gin"
)

//PeakFindOutput ...peak finding algo's output
type PeakFindOutput struct {
	Index int
	Value int
}

//PeakFind ...handles peak finding algorith
func PeakFind(c *gin.Context) {

	inputValue := c.PostForm("input")

	splitString := strings.Split(inputValue, ",")
	var input []int

	for _, s := range splitString {
		inputString := strings.Trim(s, " ")
		converted, _ := strconv.Atoi(inputString)
		input = append(input, converted)
	}

	pos := peakfinding.Peakfind(input)

	out := Ouput{
		Input: inputValue,
		Output: PeakFindOutput{
			Index: pos,
			Value: input[pos],
		},
	}
	c.JSON(200, out)
}
