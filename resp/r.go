package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type R struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	SUCCESS = 0
	ERROR   = 1
)

func Result(r *R, c *gin.Context) {
	c.JSON(http.StatusOK, r)
}

func Ok(c *gin.Context) {
	Result(&R{SUCCESS, "success", nil}, c)
}

func OkData(data interface{}, c *gin.Context) {
	Result(&R{SUCCESS, "success", data}, c)
}

func Fail(c *gin.Context) {
	Result(&R{ERROR, "error", nil}, c)
}

func FailMsg(msg string, c *gin.Context) {
	Result(&R{ERROR, msg, nil}, c)
}
