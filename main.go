package main

import (
	"agent/router"
	"fmt"
)

func main() {
	r := router.MainRouter()
	err := r.Run(":10700")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
