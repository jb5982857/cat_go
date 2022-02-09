package account

import (
	"cat/data"
	"cat/initDB"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginRequest struct {
	PhoneNum int64 `json:"phone_num"`
	code     int   `json:"code"`
}

func Login(c *gin.Context) {
	var requestData LoginRequest
	err := c.BindJSON(requestData)
	if err != nil {
		fmt.Println("login request error")
		c.JSON(http.StatusOK, data.BaseResponseData{
			Code: data.CodeParamError,
			Msg:  "参数错误",
		})
		return
	}

	initDB.Db.
}
