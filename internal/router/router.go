package router

import (
	"github.com/albingeorge/go_algos/internal/controllers"
	"github.com/gin-gonic/gin"
)

//Init .. router initialize
func Init(r *gin.Engine) {
	r.POST("/algos/peak-find", controllers.PeakFind)
}
