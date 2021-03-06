package controllers

import (
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

	input := formatInput(inputValue)

	pos := peakfinding.Peakfind(input)

	out := Output{
		Input: input,
		Output: PeakFindOutput{
			Index: pos,
			Value: input[pos],
		},
	}

	c.JSON(200, out)
}
