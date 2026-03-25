package util

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

const (
	SECRET = "fuckyoumonther"
)

func HashPassword(pwd string) (string, error) {
	ihash, err := bcrypt.GenerateFromPassword([]byte(pwd), 12)
	return string(ihash), err
}

func GenerateJWT(raw string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": raw,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})
	return token.SignedString([]byte(SECRET))
}

func CheckPwd(pwd string, hashedPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(pwd))
	return err == nil
}

func ParseJWT(token string) (string, error) {
	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	} else {
		return "", errors.New("token has problem")
	}

	t, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected Signing Method")
		}
		return []byte(SECRET), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		username, ok := claims["username"].(string)
		if !ok {
			return "", errors.New("Username claims is not string")
		}
		return username, nil
	}

	return "", err
}
