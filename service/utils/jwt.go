package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	RoleID   uint   `json:"role_id"`
	jwt.RegisteredClaims
}

func CreateJWT(tokenType, email, fullname string, roleID int) (ss string, err error) {
	mySigningKey := []byte("SecretBangetNih")
	var claims *UserClaims

	if tokenType != "access" && tokenType != "refresh" {
		err = errors.New("invalid token type")
		return "", err
	}

	if tokenType == "access" {
		claims = &UserClaims{
			email,
			fullname,
			uint(roleID),
			jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(6 * time.Hour)),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				NotBefore: jwt.NewNumericDate(time.Now()),
			},
		}
	} else if tokenType == "refresh" {
		claims = &UserClaims{
			email,
			fullname,
			uint(roleID),
			jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour)),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				NotBefore: jwt.NewNumericDate(time.Now()),
			},
		}

	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err = token.SignedString(mySigningKey)
	if err != nil {
		err = errors.New(err.Error())
		return "", err
	}

	return ss, nil
}

func ParseJWT(tokenStr string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, nil
		}

		return []byte("SecretBangetNih"), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, nil
	}

	return claims, nil
}
