package jwt

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/st0pcha/task-manager/backend/internal/config"
	"github.com/st0pcha/task-manager/backend/internal/dal"
)

func GenerateJWTToken(user *dal.User, ttl time.Duration) (string, error) {
	jwtSecret := getJWTSecretKey()

	claims := jwt.MapClaims{
		"sub": user.ID,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(ttl).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", fmt.Errorf("error signing token: %w", err)
	}
	return signedToken, nil
}

func GenerateJWTTokens(user *dal.User) (string, string, error) {
	accessToken, err := GenerateJWTToken(user, AccessTokenTTL)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := GenerateJWTToken(user, RefreshTokenTTL)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func ParseJWT(tokenStr string) (*dal.User, error) {
	jwtSecret := getJWTSecretKey()

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Println("Unexpected signing method:", token.Header["alg"])
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		if strings.Contains(err.Error(), "expired") {
			return nil, fmt.Errorf("token is expired")
		}
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	userID, ok := claims["sub"].(string)
	if !ok || userID == "" {
		return nil, fmt.Errorf("invalid or missing user ID")
	}

	user := &dal.User{}
	if err := dal.FindUserByID(user, userID).Error; err != nil {
		log.Fatalf("error fetching user: %v", err)
		return nil, fmt.Errorf("error fetching user")
	}
	// log.Println(user)

	return user, nil
}

func ValidateJWT(tokenStr string) (jwt.MapClaims, error) {
	jwtSecretKey := getJWTSecretKey()
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecretKey), nil
	})

	if err != nil {
		if err.Error() == "token is expired" {
			return nil, fmt.Errorf("token is expired")
		}
		return nil, fmt.Errorf("error parsing token: %v", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	if exp, ok := claims["exp"].(float64); ok {
		expTime := time.Unix(int64(exp), 0)
		if expTime.Before(time.Now()) {
			return nil, fmt.Errorf("token is expired")
		}
	}

	return claims, nil
}

func getJWTSecretKey() []byte {
	jwtSecret := config.JWTSecretKey
	if len(jwtSecret) == 0 {
		log.Fatal("JWT_SECRET_KEY required")
	}
	return jwtSecret
}
