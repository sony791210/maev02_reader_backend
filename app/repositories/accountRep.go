package repositories

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// custom claims
type Claims struct {
	AccountId int `json:"accountid"`
	jwt.RegisteredClaims
}

// jwt secret key
var jwtSecret = []byte("test")

func Login(account string, password string) (tokenString string, err error) {

	if account == "apple791210" && password == "qaz7963055" {
		tokenString, err = MakeToken(1)
		if err != nil {
			fmt.Println(err)
			return "", errors.New("token 組合錯誤")
		}
		return tokenString, nil
	}

	return "", errors.New("帳密錯誤")
}

func MakeToken(id int) (tokenString string, err error) {
	cliam := Claims{
		AccountId: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(3 * time.Hour * time.Duration(1))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cliam)

	tokenString, err = token.SignedString(jwtSecret)
	return tokenString, err
}
