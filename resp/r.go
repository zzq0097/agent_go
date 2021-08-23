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
	OK    = 0
	ERROR = 1
)

func Result(r *R, c *gin.Context) {
	c.JSON(http.StatusOK, r)
}

func Success(c *gin.Context) {
	Result(&R{OK, "success", nil}, c)
}

func SuccessData(data interface{}, c *gin.Context) {
	Result(&R{OK, "success", data}, c)
}

func Error(c *gin.Context) {
	Result(&R{ERROR, "error", nil}, c)
}

func ErrorMsg(msg string, c *gin.Context) {
	Result(&R{ERROR, msg, nil}, c)
}
