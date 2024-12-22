package checkers

import (
	"errors"
	"regexp"
)

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	if len(password) > 32 {
		return errors.New("password must be at most 32 characters long")
	}
	if !regexp.MustCompile(`[A-Z]+`).MatchString(password) {
		return errors.New("password must contain at least one uppercase letter")
	}
	if !regexp.MustCompile(`[a-z]+`).MatchString(password) {
		return errors.New("password must contain at least one lowercase letter")
	}
	return nil
}

func ValidatePasswordsMatch(password, repeatPassword string) error {
	if password != repeatPassword {
		return errors.New("passwords don't match")
	}
	return nil
}
