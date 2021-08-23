package router

import (
	"agent/api"
	"github.com/gin-gonic/gin"
)

func ConfRouter(r *gin.Engine) {
	conf := r.Group("/conf")
	{
		conf.GET("/allFile", api.AllFile)
		conf.GET("/replace", api.Replace)
		conf.GET("/rollback", api.Rollback)
	}
}
