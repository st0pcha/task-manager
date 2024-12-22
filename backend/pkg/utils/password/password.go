package password

import (
	"golang.org/x/crypto/bcrypt"
)

func Hash(raw string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(raw), 10)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

func Verify(hash, raw string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(raw))
	return err == nil
}
