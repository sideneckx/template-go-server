package server_jwt

import (
	"errors"
	"log"

	"github.com/golang-jwt/jwt/v4"
)

var secret = []byte("phat'll be the king")

func GenToken(data map[string]interface{}) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(data))
	tokenString, err := token.SignedString(secret)
	if err != nil {
		log.Fatalln(err)
	}
	return tokenString
}

func GenTokenUserId(userId string) string {
	data := map[string]interface{}{"userId": userId}
	return GenToken(data)
}

func DecodeToken(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("invalid token")
	}
}

func DecodeTokenUserId(tokenString string) (string, error) {
	claims, err := DecodeToken(tokenString)
	if err != nil {
		return "", err
	}
	userId := claims["userId"].(string)
	return userId, nil
}
