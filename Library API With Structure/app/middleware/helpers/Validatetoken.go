package helpers

import (
	"github.com/dgrijalva/jwt-go"
)

func Validatetoken(tokenString string) (bool, string) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, nil
		}
		return []byte("secret"), nil
	})
	if err != nil {
		return false, err.Error()
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["authorized"] == true {
			return true, ""
		}
		return false, "Token not authorized"
	}
	return false, "Token not valid"
}