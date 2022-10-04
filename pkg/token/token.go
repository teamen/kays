package token

import (
	"fmt"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Config struct {
	signingSecret string
}

var (
	config = Config{"Rtg8BPKNEf2mB4mgvKONGPZZQSaJWNLijxR42qRgq0iBb5"}
	once   sync.Once
)

func GetSignSecret() string {
	return config.signingSecret
}

func Sign(id int, username string) (string, error) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       id,
		"username": username,
		"nbf":      time.Now().Unix(),
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(2 * time.Hour).Unix(),
	}).SignedString([]byte(config.signingSecret))

	if err != nil {
		return "", err
	}

	return token, nil
}

func Parse(tokenString string) (id int, username string, err error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.signingSecret), nil
	})
	if err != nil {
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id = int(claims["id"].(float64))
		username = claims["username"].(string)

		return
	}

	err = fmt.Errorf("unable to verify token")
	return
}

func Init(signingSecret string) {
	once.Do(func() {
		if config.signingSecret != "" {
			config.signingSecret = signingSecret
		}
	})
}
