package main

import (
	"github.com/albingeorge/go_algos/internal/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("cmd/templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(
			200,
			"index.html",
			gin.H{
				"title": "Home Page",
			},
		)
	})

	router.Init(r)

	r.Run()
}
