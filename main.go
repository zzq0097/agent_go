package main

import (
	"agent/router"
)

func main() {
	r := router.MainRouter()
	r.Run(":8877")
}
