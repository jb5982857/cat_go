package account

import (
	"cat/data"
	"cat/initDB"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginRequest struct {
	PhoneNum int `json:"phone_num"`
	Code     int `json:"code"`
}

func Login(c *gin.Context) {
	requestData := LoginRequest{}
	err := c.BindJSON(&requestData)
	if err != nil {
		fmt.Println("login request error " + err.Error())
		c.JSON(http.StatusOK, data.BaseResponseData{
			Code: data.CodeParamError,
			Msg:  "参数错误",
		})
		return
	}

	_, e := initDB.Db.Exec("insert into account(phone) values (?);", requestData.PhoneNum)
	if e != nil {
		fmt.Println("save account error " + e.Error())
	}
}
