package main

import (
	account "cat/api"
	"cat/data"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST(data.Login, account.Login)
	err := router.Run(":8090")
	if err != nil {
		fmt.Printf("router start err %s", err.Error())
		return
	}
}
