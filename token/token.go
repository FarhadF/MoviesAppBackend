package token

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateToken(role string) (string, error) {
	mySigningKey := "Super_Dup3r_S3cret"
	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)
	// Set some claims
	token.Claims = jwt.MapClaims{
		"exp":  time.Now().Add(time.Microsecond * time.Duration(1)).Unix(),
		"iat":  time.Now().Unix(),
		"role": role,
	}

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(mySigningKey))
	fmt.Println(tokenString)
	return tokenString, err

}
