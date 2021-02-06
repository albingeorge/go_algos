package controllers

import (
	"errors"
	"strconv"

	"github.com/albingeorge/go_algos/internal/algos/mit_intro_to_algos/search"
	"github.com/albingeorge/go_algos/internal/constants"
	"github.com/gin-gonic/gin"
)

//SearchOutput ... search output structure
type SearchOutput struct {
	Index int    `json:"index"`
	Value int    `json:"value"`
	Error string `json:"error"`
}

//Search .. search algorithms
func Search(c *gin.Context) {
	var pos int
	var err error
	out := Output{}

	inputString := c.PostForm("input")
	algo := c.PostForm("algorithm")

	x, err := strconv.Atoi(c.PostForm("search"))
	if err != nil {
		out.Output = SearchOutput{
			Error: "expected integer value to be passed",
		}
	}

	input := formatInput(inputString)

	out.Input = input

	switch algo {
	case constants.BinarySearch:
		pos, err = search.Binary(input, x)
	default:
		err = errors.New("algorithm not found: " + algo)
	}

	if err != nil {
		out.Output = SearchOutput{
			Error: err.Error(),
		}
		c.JSON(200, out)
	}

	out.Output = SearchOutput{
		Index: pos,
		Value: x,
	}

	c.JSON(200, out)
}
