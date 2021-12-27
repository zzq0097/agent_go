package router

import (
	"agent/api"
	"github.com/gin-gonic/gin"
)

func ParameterRouter(r *gin.Engine) {
	parameter := r.Group("/parameter")
	{
		parameter.GET("/get", api.GetBasicParameter)
		parameter.POST("/set", api.SetBasicParameter)
	}
}
