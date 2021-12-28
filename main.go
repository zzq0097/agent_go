package main

import (
	"agent/config"
	"agent/crontab"
	"agent/router"
	"agent/util"
)

func main() {
	util.InitDB()
	config.Init()
	crontab.RunCron()

	r := router.MainRouter()
	err := r.Run(":10700")
	if err != nil {
		panic(err)
		return
	}
}
