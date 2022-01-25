package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	ISSUER            = "github.com/kazmerdome/go-graphql-starter"
	ERR_INVALID_TOKEN = "invalid stored token"
)

/**
 * JWT
 */

// CustomClaims is a custom claim for storing custom map[string]string data
// It extends the jwt.StandardClaims claims with a data field
type CustomClaims struct {
	Data map[string]string `json:"data"`
	jwt.StandardClaims
}

func GenerateJWTToken(data map[string]string, sessionSecret string, sessionExpiration int32) (string, error) {
	if data == nil {
		return "", errors.New("data is required")
	}

	expireToken := time.Now().Add(time.Hour * time.Duration(sessionExpiration)).Unix()
	claims := &CustomClaims{
		data,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			IssuedAt:  time.Now().Unix(),
			Id:        primitive.NewObjectID().Hex(),
			Issuer:    ISSUER,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(sessionSecret))
	if err != nil {
		return "", err
	}
	return tokenString, err
}

func VerifyJWTToken(tokenString string, sessionSecret string) (map[string]string, error) {
	data := make(map[string]string)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(sessionSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		dataMapStringInterface, _ok := claims["data"].(map[string]interface{})
		if !_ok {
			return nil, errors.New(ERR_INVALID_TOKEN)
		}
		// if dataMapStringInterface == nil {
		// 	return nil, errors.New(ERR_INVALID_TOKEN)
		// }
		for k, v := range dataMapStringInterface {
			val, __ok := v.(string)
			if !__ok {
				return nil, errors.New(ERR_INVALID_TOKEN)
			}
			data[k] = val
		}
	} else {
		return nil, errors.New(ERR_INVALID_TOKEN)
	}

	return data, nil
}
