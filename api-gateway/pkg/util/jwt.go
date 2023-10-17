package util

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"time"
)

var jwtSecret []byte

var jwtInstance *jwt.Token

type Claims struct {
	jwt.StandardClaims
	data interface{} `json:"data"`
}

func init() {
	token := jwt.New(jwt.SigningMethodHS256)
	jwtInstance = token
	//jwtSecret = config.JwtSecret
	jwtSecret = []byte(viper.GetString("server.jwtSecret"))
}

func GenerateToken(data interface{}) (tokenStr string, err error) {
	expireTime := time.Now().Add(time.Hour * 24)

	claim := Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "",
		},
		data: data,
	}

	jwtInstance.Claims = claim

	tokenStr, err = jwtInstance.SignedString(jwtSecret)

	return
}

func ParseToken(tokenStr string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if data, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
		return data, nil
	} else {
		return nil, errors.New("invalid token")
	}
}
