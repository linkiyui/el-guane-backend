package auth_token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateRefreshToken(user_id string) (string, error) {
	exp := 12
	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user_id,
		"exp":     time.Now().Add(time.Hour * 24 * time.Duration(exp)).Unix(),
		"iat":     time.Now().Unix(),
		"sub":     "refresh",
	})
	secret := "supersecret"
	return tkn.SignedString([]byte(secret))
}

func ValidateRefreshToken(token string) (*LoginClaims, error) {

	secret := "supersecret"

	jwtKeyFunc := func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secret), nil
	}

	tkn, err := jwt.ParseWithClaims(token, &LoginClaims{}, jwtKeyFunc, jwt.WithExpirationRequired(), jwt.WithSubject("refresh"))

	if err != nil {
		return nil, err
	}

	claims, ok := tkn.Claims.(*LoginClaims)
	if !ok {
		return nil, fmt.Errorf("can't parse claims: %v", tkn.Claims)
	}

	return claims, nil
}
