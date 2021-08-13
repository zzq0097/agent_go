package route

import (
	"agent/consts"
	"agent/resp"
	"agent/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CommandRoute(r *gin.Engine) {
	command := r.Group("/command")
	{
		command.GET("/status", func(c *gin.Context) {
			c.JSON(http.StatusOK, resp.Success())
		})
		command.GET("/version", func(c *gin.Context) {
			c.JSON(http.StatusOK, resp.SuccessData(service.NgxVersion(consts.NgxDefCmd)))
		})
		command.GET("/start", func(c *gin.Context) {
			c.JSON(http.StatusOK, resp.Success())
		})
		command.GET("/stop", func(c *gin.Context) {
			c.JSON(http.StatusOK, resp.Success())
		})
		command.GET("/reload", func(c *gin.Context) {
			c.JSON(http.StatusOK, resp.Success())
		})
		command.GET("/check", func(c *gin.Context) {
			c.JSON(http.StatusOK, resp.Success())
		})
	}
}
