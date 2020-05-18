package utils

import (
	"fmt"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	jwtKey = "123123"
	token  = ""
)

func TestCreateToken(t *testing.T) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		ID:    1,
		Email: "dev@yopmail.com",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	tokenTemp, err := CreateToken(claims, jwtKey)
	if err != nil {
		t.Error("Error:", err)
	} else {
		token = tokenTemp
		fmt.Println(tokenTemp)
	}
}

func TestDecodeToken(t *testing.T) {

	claims := &Claims{}
	err := DecodeToken(token, claims, jwtKey)
	if err != nil {
		t.Error("Error:", err)
	} else {
		if claims.Email != "dev@yopmail.com" {
			t.Error("Claim email must be dev@yopmail.com")
		}
	}
}
