package router

import (
	"agent/api"
	"github.com/gin-gonic/gin"
)

func LogRouter(r *gin.Engine) {
	log := r.Group("/log")
	{
		log.GET("/ReverseLines", api.LogLines)
	}
}
