package checkers

import (
	"errors"
	"regexp"
)

func ValidateEmail(email string) error {
	if !regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,4}$`).MatchString(email) {
		return errors.New("invalid email format")
	}
	return nil
}
