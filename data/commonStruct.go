package data

import "github.com/dgrijalva/jwt-go"

type BaseResponseData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type AccountJwtData struct {
	AccountId int `json:"account_id"`
	jwt.StandardClaims
}
