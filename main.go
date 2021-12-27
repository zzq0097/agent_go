package main

import (
	"agent/config"
	"agent/router"
	"agent/util"
	"fmt"
)

func main() {
	util.InitDB()
	config.Init()

	r := router.MainRouter()
	err := r.Run(":10700")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
