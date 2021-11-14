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
	OK   = 0
	FAIL = 1
)

func Result(r *R, c *gin.Context) {
	c.JSON(http.StatusOK, r)
}

func Ok(c *gin.Context) {
	Result(&R{OK, "success", nil}, c)
}

func OkData(data interface{}, c *gin.Context) {
	Result(&R{OK, "success", data}, c)
}

func Fail(c *gin.Context) {
	Result(&R{FAIL, "error", nil}, c)
}

func FailMsg(msg string, c *gin.Context) {
	Result(&R{FAIL, msg, nil}, c)
}
