package token

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"strings"
	"time"
)

type errOut struct {
	Error string `json:error`
}

const mySigningKey = "Super_Dup3r_S3cret"

func GenerateToken(role string) (string, error) {
	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)
	// Set some claims
	token.Claims = jwt.MapClaims{
		"exp":  time.Now().Add(time.Hour * time.Duration(1)).Unix(),
		"iat":  time.Now().Unix(),
		"role": role,
	}

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(mySigningKey))
	fmt.Println(tokenString)
	return tokenString, err
}

func ParseToken(myToken string) (bool, error) {
	token, err := jwt.Parse(myToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(mySigningKey), nil
	})
	//fmt.Println(token.Claims)
	if err == nil && token.Valid {
		return true, nil
	} else {
		return false, err
	}
}

func ExtractToken(r *http.Request) (string, error) {
	auth := r.Header.Get("Authorization")
	split := strings.Split(auth, " ")
	if split[0] != "Bearer" || split[1] == "" {

		return "", errors.New("Malformed Auth Header")
	} else {
		return split[1], nil
	}

}

func TokenHandler(w http.ResponseWriter, r *http.Request) bool {

	authToken, err := ExtractToken(r)
	if err != nil {
		errout := new(errOut)
		errout.Error = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(errout); err != nil {
			log.Panic("Error EncodingJson in TokenHandler", err)
		}
		return false
	} else {
		tokenStatus, err := ParseToken(authToken)
		//fmt.Println("tokenStatus: ", tokenStatus, "err: ", err.Error())
		if err != nil || tokenStatus == false {
			errout := new(errOut)
			errout.Error = err.Error()
			w.WriteHeader(http.StatusForbidden)
			if err := json.NewEncoder(w).Encode(errout); err != nil {
				log.Panic("Error EncodingJson in TokenHandler", err)
			}
			log.Println("token status err: ", err)
			return false

		} else {

			//w.Header().Set("Access-Control-Allow-Origin", "*")
			return true

		}
	}

}
