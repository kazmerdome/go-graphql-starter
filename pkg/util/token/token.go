package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	ISSUER            = "github.com/kazmerdome/go-graphql-starter"
	ERR_INVALID_TOKEN = "invalid stored token"
)

/**
 * JWT
 */
func GenerateJWTToken(ID string, sessionSecret string, sessionExpiration int32) (string, error) {
	expireToken := time.Now().Add(time.Hour * time.Duration(sessionExpiration)).Unix()
	claims := &jwt.StandardClaims{
		ExpiresAt: expireToken,
		IssuedAt:  time.Now().Unix(),
		Id:        ID,
		Issuer:    ISSUER,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(sessionSecret))
	if err != nil {
		return "", err
	}
	return tokenString, err
}

func VerifyJWTToken(tokenString string, sessionSecret string) (string, error) {
	var id string

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(sessionSecret), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["jti"].(string) == "" {
			return "", errors.New(ERR_INVALID_TOKEN)
		}
		id = claims["jti"].(string)
	} else {
		return "", errors.New(ERR_INVALID_TOKEN)
	}

	return id, nil
}
