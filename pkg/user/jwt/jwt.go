package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"login-task/pkg/configs"
	"login-task/pkg/user/storage"
	"strings"
	"time"
)

type Claims struct {
	jwt.StandardClaims
	ID string `json:"id"`
}

var config = configs.GetConfig()

func CreateAccessToken(id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(config.AccessTTL).Unix(),
		},
		ID: id,
	})
	return token.SignedString([]byte(config.JWTSigningKey))
}

func CreateRefreshToken(id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(config.RefreshTTL).Unix(),
		},
		ID: id,
	})
	res, _ := token.SignedString([]byte(config.JWTSigningKey))
	if err := storage.AddRefreshToken(res, id); err != nil {
		return "", err
	}
	return res, nil
}

func ParseToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(config.JWTSigningKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return "", nil
	}
	return claims.ID, nil
}

func TokenCheck(header string) (string, error) {
	if header == "" {
		return "", errors.New("error")
	}

	headerSplit := strings.Split(header, " ")
	if len(headerSplit) != 2 {
		return "", errors.New("error")
	}

	userID, err := ParseToken(headerSplit[1])
	if err != nil {
		return "", errors.New("error")
	}

	return userID, nil
}
