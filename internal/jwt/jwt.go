package jwt

import (
	"context"
	"time"

	"github.com/blawhi2435/shanjuku-backend/environment"
	"github.com/blawhi2435/shanjuku-backend/internal/cerror"
	"github.com/blawhi2435/shanjuku-backend/internal/contextkey"
	"github.com/dgrijalva/jwt-go"
)

const (
	issuer = "GiftForm69King"
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
			Issuer:    issuer,
			Audience:  data.Account,
		},
	}
	// Choose specific algorithm
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// Choose specific Signature
	return token.SignedString([]byte(environment.Setting.Auth.JWTSecret))
}

func parseToken(tokenString string) (*MyClaims, error) {
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
	return nil, cerror.ErrTokenInvalid
}

func ValidateTokenAndGetPayload(ctx context.Context) (PayloadData, bool) {
	tokenCtx := ctx.Value(contextkey.TokenCtxKey)
	if tokenCtx == nil || tokenCtx == "" {
		return PayloadData{}, false
	}

	token, ok := tokenCtx.(string)
	if !ok {
		return PayloadData{}, false
	}

	mc, err := parseToken(token)
	if err != nil {
		return PayloadData{}, false
	}

	if mc.Issuer != issuer {
		return PayloadData{}, false
	}

	return mc.PayloadData, true
}
