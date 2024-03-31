package util

import "golang.org/x/crypto/bcrypt"

func CompareHash(input, stored string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(stored), []byte(input))
	return err == nil
}

func GenerateHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
