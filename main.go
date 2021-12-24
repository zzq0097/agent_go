package main

import (
	"agent/init"
	"agent/router"
	"fmt"
)

func main() {
	init.Init()

	r := router.MainRouter()
	err := r.Run(":10700")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
