package router

import (
	"agent/api"
	"github.com/gin-gonic/gin"
)

func CodeRouter(r *gin.Engine) {
	code := r.Group("/code")
	{
		code.GET("/get", api.Get)
	}
}
