package jwt

import (
	"errors"
	"time"

	"github.com/blawhi2435/shanjuku-backend/environment"
	"github.com/dgrijalva/jwt-go"
)

type PayloadData struct {
	UserID  int64  `json:"user_id"`
	Account string `json:"account"`
}

type MyClaims struct {
	PayloadData
	jwt.StandardClaims
}

func GenToken(data PayloadData) (string, error) {
	c := MyClaims{
		PayloadData: data,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(30 * 24 * time.Hour).Unix(),
			Issuer:    "GiftForm69King",
			Audience:  data.Account,
		},
	}
	// Choose specific algorithm
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// Choose specific Signature
	return token.SignedString([]byte(environment.Setting.Auth.JWTSecret))
}

func ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(environment.Setting.Auth.JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}
	// Valid token
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
