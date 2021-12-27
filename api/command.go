package api

import (
	"agent/consts"
	"agent/dao"
	"agent/resp"
	"agent/service"
	"github.com/gin-gonic/gin"
)

func NgxStatus(c *gin.Context) {
	binFile, binErr := dao.GetParameter(consts.BinName)
	confFile, confErr := dao.GetParameter(consts.ConfName)
	if binErr != nil || confErr != nil {
		resp.FailMsg("参数获取失败", c)
	}
	ngxStatus := service.NgxStatus(binFile, confFile)
	resp.OkData(ngxStatus, c)
}

func NgxStart(c *gin.Context) {
	binFile, binErr := dao.GetParameter(consts.BinName)
	confFile, confErr := dao.GetParameter(consts.ConfName)
	if binErr != nil || confErr != nil {
		resp.FailMsg("参数获取失败", c)
	}
	ngxStatus := service.NgxStart(binFile, confFile)
	resp.OkData(ngxStatus, c)
}

func NgxReload(c *gin.Context) {
	binFile, binErr := dao.GetParameter(consts.BinName)
	confFile, confErr := dao.GetParameter(consts.ConfName)
	if binErr != nil || confErr != nil {
		resp.FailMsg("参数获取失败", c)
	}
	ngxStatus := service.NgxReload(binFile, confFile)
	resp.OkData(ngxStatus, c)
}

func NgxStop(c *gin.Context) {
	binFile, binErr := dao.GetParameter(consts.BinName)
	confFile, confErr := dao.GetParameter(consts.ConfName)
	if binErr != nil || confErr != nil {
		resp.FailMsg("参数获取失败", c)
	}
	ngxStatus := service.NgxStop(binFile, confFile)
	resp.OkData(ngxStatus, c)
}

func NgxV(c *gin.Context) {
	binFile, binErr := dao.GetParameter(consts.BinName)
	if binErr != nil {
		resp.FailMsg("参数获取失败", c)
	}
	resp.OkData(service.NgxV(binFile), c)
}

func NgxVersion(c *gin.Context) {
	binFile, binErr := dao.GetParameter(consts.BinName)
	if binErr != nil {
		resp.FailMsg("参数获取失败", c)
	}
	resp.OkData(service.NgxVersion(service.NgxV(binFile)), c)
}

func NgxCfgArgs(c *gin.Context) {
	binFile, binErr := dao.GetParameter(consts.BinName)
	if binErr != nil {
		resp.FailMsg("参数获取失败", c)
	}
	resp.OkData(service.NgxCfgArgs(service.NgxV(binFile)), c)
}

func NgxTConf(c *gin.Context) {
	binFile, binErr := dao.GetParameter(consts.BinName)
	confFile, confErr := dao.GetParameter(consts.ConfName)
	if binErr != nil || confErr != nil {
		resp.FailMsg("参数获取失败", c)
	}
	resp.OkData(service.NgxTConf(binFile, confFile), c)
}

func CheckPortUsed(c *gin.Context) {
	resp.OkData(service.CheckPortUsed(), c)
}

func Os(c *gin.Context) {
	resp.OkData(service.Os(), c)
}
