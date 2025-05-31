package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(passwordString string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword(
		[]byte(passwordString),
		14,
	)

	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

func CheckHashPassword(hashedPassword string, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(password),
	)

	if err != nil {
		return false, bcrypt.ErrMismatchedHashAndPassword
	}

	return true, nil
}
