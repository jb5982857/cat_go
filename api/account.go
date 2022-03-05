package account

import (
	"cat/data"
	loginHandler "cat/handler"
	"cat/initDB"
	"cat/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AccountLogin(c *gin.Context) {
	requestData := loginHandler.AccountRequest{}
	err := c.BindJSON(&requestData)
	if err != nil {
		fmt.Println("account request error " + err.Error())
		c.JSON(http.StatusOK, data.BaseResponseData{
			Code: data.CodeParamError,
			Msg:  "参数错误",
		})
		return
	}

	body, token := requestData.AccountLogin()
	c.Header(data.KeyAccountToken, token)
	c.JSON(http.StatusOK, body)
}

func Register(c *gin.Context) {
	requestData := loginHandler.AccountRequest{}
	err := c.BindJSON(&requestData)
	if err != nil {
		fmt.Println("register request error " + err.Error())
		c.JSON(http.StatusOK, data.BaseResponseData{
			Code: data.CodeParamError,
			Msg:  "参数错误",
		})
		return
	}

	c.JSON(http.StatusOK, requestData.Register())
}

func TestToken(c *gin.Context) {
	token := c.GetHeader(data.KeyAccountToken)
	if token != "" {
		accountData := utils.ParseAccountIdToken(token)
		fmt.Println("token test ", accountData.AccountId, accountData.ExpiresAt)
	}
}

func PhoneLogin(c *gin.Context) {
	requestData := loginHandler.PhoneLoginRequest{}
	err := c.BindJSON(&requestData)
	if err != nil {
		fmt.Println("login request error " + err.Error())
		c.JSON(http.StatusOK, data.BaseResponseData{
			Code: data.CodeParamError,
			Msg:  "参数错误",
		})
		return
	}

	phone, err := initDB.Db.Query("SELECT phone from account where phone=?", requestData.PhoneNum)
	if err != nil {
		fmt.Println(phone, " query db error")
		return
	}

	if phone != nil {
		fmt.Println("phone is not null, ", phone)
		_, err := initDB.Db.Exec("UPDATE account set last_login_time=? where phone=?", utils.GetCurrentDataTimeToDb(), requestData.PhoneNum)
		if err != nil {
			fmt.Println("update error ", err.Error())
			return
		}
	} else {
		_, e := initDB.Db.Exec("insert into account(phone , last_login_time) values (?,?);", requestData.PhoneNum, utils.GetCurrentDataTimeToDb())
		if e != nil {
			fmt.Println("save account error " + e.Error())
		}
	}

}
