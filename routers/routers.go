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
		s.POST(data.PhoneLogin, account.PhoneLogin)
		s.POST(data.AccountLogin, account.AccountLogin)
		s.POST(data.Register, account.Register)
	}
	return r
}
