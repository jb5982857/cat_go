package routers

import (
	account "cat/api"
	"cat/data"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	s := r.Group(data.Group)
	{
		s.POST(data.Login, account.Login)
	}
	return r
}
