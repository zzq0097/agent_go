package main

import (
	"agent/route"
)

func main() {
	r := route.MainRoute()
	r.Run(":8877")
}
