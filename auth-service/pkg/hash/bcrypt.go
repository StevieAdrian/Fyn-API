package hash

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password *string) *string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	hashed := string(bytes)
	return &hashed
}

func VerifyPassword(hashed, plain string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	return err == nil, err
}
