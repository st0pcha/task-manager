package checkers

import (
	"fmt"
	"regexp"
)

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}
	if len(password) > 32 {
		return fmt.Errorf("password must be at most 32 characters long")
	}
	if !regexp.MustCompile(`[A-Z]+`).MatchString(password) {
		return fmt.Errorf("password must contain at least one uppercase letter")
	}
	if !regexp.MustCompile(`[a-z]+`).MatchString(password) {
		return fmt.Errorf("password must contain at least one lowercase letter")
	}
	return nil
}

func ValidatePasswordsMatch(password, repeatPassword string) error {
	if password != repeatPassword {
		return fmt.Errorf("passwords don't match")
	}
	return nil
}
