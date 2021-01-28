package controllers

import (
	"fmt"

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

	out := Ouput{
		Input: inputValue,
		Output: PeakFindOutput{
			Index: 0,
			Value: 2,
		},
	}
	fmt.Println(out)
	c.JSON(200, out)
}
