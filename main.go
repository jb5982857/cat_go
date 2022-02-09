package main

import (
	"cat/routers"
	"fmt"
)

func main() {
	r := routers.Router()
	err := r.Run(":8090")
	if err != nil {
		fmt.Printf("routers start err %s", err.Error())
		panic("start error")
	}
}
