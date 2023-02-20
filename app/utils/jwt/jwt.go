package jwt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWT struct {
	Username string
	Password string
	Exp      int64
	Token    string
}

// CreateToken creates a new JWT token
func (j *JWT) CreateToken() (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user"] = j.Username
	claims["password"] = j.Password
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires after 24 hours

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("secret"))
}

// ValidateToken checks if the token is valid
func (j *JWT) ValidateToken() (bool, error) {
	token, err := jwt.Parse(j.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})

	if err != nil {
		return false, err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true, nil
	}

	return false, nil
}
