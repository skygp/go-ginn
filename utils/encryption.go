package utils

import "golang.org/x/crypto/bcrypt"

func CompareHashAndPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func GenerateFromPassword(password string) (hashedPassword string, err error) {
	var hashPwd []byte
	hashPwd,err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashPwd),err
}