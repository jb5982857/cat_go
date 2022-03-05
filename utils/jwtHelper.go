package utils

import (
	"cat/data"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var Secret = []byte("fdsjakfdsajklfsda")

const TokenExpireDuration = 2 * time.Hour

func GetAccountIdJwtToken(data *data.AccountJwtData) string {
	data.ExpiresAt = time.Now().Add(TokenExpireDuration).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	signedString, err := token.SignedString(Secret)
	if err != nil {
		fmt.Println("GetAccountIdJwtToken Error ", err.Error())
		return ""
	}

	return signedString
}

func ParseAccountIdToken(tokenStr string) *data.AccountJwtData {
	token, err := jwt.ParseWithClaims(tokenStr, &data.AccountJwtData{}, func(tk *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if err != nil {
		return nil
	}
	if resultData, ok := token.Claims.(*data.AccountJwtData); ok && token.Valid {
		return resultData
	}
	return nil
}
