package loginHandler

import (
	"cat/data"
	"cat/initDB"
	dataHelper "cat/utils"
	"fmt"
	"strings"
)

type PhoneLoginRequest struct {
	PhoneNum int `json:"phone_num"`
	Code     int `json:"code"`
}

type AccountRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(account *AccountRequest) data.BaseResponseData {
	count := 0
	queryErr := initDB.Db.QueryRow("SELECT COUNT(*) from account where username=?", account.Username).Scan(&count)
	if queryErr != nil {
		fmt.Println("register request error " + queryErr.Error())
		return data.BaseResponseData{
			Code: data.CodeDbFailed,
			Msg:  "注册失败",
		}
	}

	if count != 0 {
		return data.BaseResponseData{
			Code: data.CodeFailed,
			Msg:  "账号已存在",
		}
	}

	_, e := initDB.Db.Exec("insert into account(username , password,last_login_time) values (?,?,?);", account.Username, strings.ToLower(account.Password), dataHelper.GetCurrentDataTimeToDb())

	if e != nil {
		fmt.Println("register request error " + e.Error())
		return data.BaseResponseData{
			Code: data.CodeDbFailed,
			Msg:  "注册失败",
		}
	}

	return data.BaseResponseData{
		Code: data.CodeSuccess,
		Msg:  "注册成功",
	}
}

func AccountLogin(account *AccountRequest) data.BaseResponseData {
	dbAccount := ""
	err := initDB.Db.QueryRow("SELECT username from account where username=? and password=? LIMIT 1", account.Username, strings.ToLower(account.Password)).Scan(&dbAccount)
	if err != nil {
		return data.BaseResponseData{
			Code: data.CodeDbFailed,
			Msg:  "登录失败",
		}
	}

	if dbAccount != "" {
		initDB.Db.Exec("UPDATE account set last_login_time=?", dataHelper.GetCurrentDataTimeToDb())
		return data.BaseResponseData{
			Code: data.CodeSuccess,
			Msg:  "登录成功",
		}
	}

	return data.BaseResponseData{
		Code: data.CodeFailed,
		Msg:  "账号密码错误",
	}
}
