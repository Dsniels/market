package services

import (
	"fmt"
	"time"

	"github.com/dsniels/market/core/types"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwt.RegisteredClaims
}

type Token struct {
}

func getKey() string {
	return "this is a super secret key nobody has to know about it"
}

func (t *Token) GenerateToken(user *types.User) (*string, error) {
	claims := getClaims()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	str, err := token.SignedString([]byte(getKey()))
	if err != nil {
		return nil, err
	}

	return &str, nil
}

func (t *Token) ValidateToken(tokenEncr string) (bool, *Claims) {
	token, err := jwt.ParseWithClaims(tokenEncr, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Invalid token")
		}

		return []byte(getKey()), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return false, nil
	}

	if !token.Valid {
		return false, nil
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return false, nil
	}
	return true, claims
}

func getClaims() *Claims {

	return &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "this",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

}

func NewToken() *Token {
	return &Token{}
}
