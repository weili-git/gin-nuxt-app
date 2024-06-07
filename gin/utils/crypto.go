package utils

import "golang.org/x/crypto/bcrypt"

func Encrypt(text string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
