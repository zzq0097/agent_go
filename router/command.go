package router

import (
	"agent/api"
	"github.com/gin-gonic/gin"
)

func CommandRouter(r *gin.Engine) {
	command := r.Group("/command")
	{
		command.GET("/status", api.NgxStatus)
		command.GET("/start", api.NgxStart)
		command.GET("/stop", api.NgxStop)
		command.GET("/reload", api.NgxReload)
		command.GET("/ngxV", api.NgxV)
		command.GET("/version", api.NgxVersion)
		command.GET("/cfgArgs", api.NgxCfgArgs)
	}
}
