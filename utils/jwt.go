package utils

import "github.com/golang-jwt/jwt/v5"

var SecretKey = "SECRETKEY"

func GenerateToken(claims *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	webtoken, err := token.SignedString([]byte(SecretKey))

	if err != nil {
		return "", err
	}
	return webtoken, nil
}
