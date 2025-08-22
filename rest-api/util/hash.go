package util

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bycrypted, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bycrypted), err
}

func CompareHashAndPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
