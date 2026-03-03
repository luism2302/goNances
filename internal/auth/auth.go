package auth

import (
	"errors"

	"github.com/alexedwards/argon2id"
)

func HashPassword(password string) (string, error) {
	hashed, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return "", errors.New("Couldn't hash password")
	}
	return hashed, nil
}

func CheckHashedPassword(password, hashedPassword string) (bool, error) {
	match, _, err := argon2id.CheckHash(password, hashedPassword)
	if err != nil {
		return false, errors.New("Couldn't check password against hash")
	}
	return match, nil
}
