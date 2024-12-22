package checkers

import (
	"fmt"
	"regexp"
)

func ValidateEmail(email string) error {
	if !regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,4}$`).MatchString(email) {
		return fmt.Errorf("invalid email format")
	}
	return nil
}
