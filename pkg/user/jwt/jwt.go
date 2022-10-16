package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"strings"
	"time"
)

type Claims struct {
	jwt.StandardClaims
	Id string `json:"id"`
}

func CreateToken(id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
		},
		Id: id,
	})
	return token.SignedString([]byte("sadfddKDAJFljaskd7usf"))
}

func ParseToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte("sadfddKDAJFljaskd7usf"), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return "", nil
	}
	return claims.Id, nil
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
