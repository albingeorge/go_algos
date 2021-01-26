package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("cmd/templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(
			200,
			"index.html",
			gin.H{
				"title": "Home Page",
			},
		)
	})
	router.Run()
}
