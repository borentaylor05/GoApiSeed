package api

import (
	"fmt"
	"time"

	"github.com/borentaylor05/streamline/util"
	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateToken(username, userType string) (string, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"type":     userType,
		"nbf":      time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	secret, err := getSecret()
	if err != nil {
		return "", err
	}
	tokenString, err := token.SignedString([]byte(secret))

	return tokenString, err
}

func ParseClaims(tokenString string) (jwt.MapClaims, error) {
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		secret, err := getSecret()
		if err != nil {
			return nil, err
		}
		return []byte(secret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("Error parsing token")
}

func getSecret() (string, error) {
	config, err := util.ReadConfig()
	if err != nil {
		return "", err
	}
	return config.Server.HmacSecret, nil
}
