package main

import (
	"fmt"

	"scheduler/methods"
)

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	methods.HandleRequest()
}
