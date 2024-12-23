package jwt

import "time"

var (
	AccessTokenTTL  = time.Hour * 1
	RefreshTokenTTL = time.Hour * 24 * 7
)

const (
	AccessToken  = "accessToken"
	RefreshToken = "refreshToken"
)
