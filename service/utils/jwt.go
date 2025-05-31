package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	RoleID   uint   `json:"role_id"`
	jwt.RegisteredClaims
}

func CreateJWT(email string, fullname string, roleID int) (string, error) {
	mySigningKey := []byte("SecretBangetNih")
	claims := UserClaims{
		email,
		fullname,
		uint(roleID),
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}

	return ss, err
}

func VerifyJWT(tokenStr string) (*UserClaims, error) {
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
