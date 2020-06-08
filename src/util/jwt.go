package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Claims struct
type Claims struct {
	ID    uint64 `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

// CreateToken func
func CreateToken(claims *Claims, jwtKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtKey))
}

// CreateRefreshToken func
func CreateRefreshToken(jwtKey string) (string, error) {
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = 1
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	return refreshToken.SignedString([]byte(jwtKey))
}

// CreateAuthTokenPair func
func CreateAuthTokenPair(email string, id uint64, role string, jwtKey string, expMinute int) (string, string, error) {
	expirationTime := time.Now().Add(time.Duration(expMinute) * time.Minute)
	claims := &Claims{
		ID:    id,
		Email: email,
		Role:  role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	authToken, err := CreateToken(claims, jwtKey)
	if err != nil {
		return "", "", err
	}
	refreshToken, err := CreateRefreshToken(jwtKey)
	if err != nil {
		return "", "", err
	}
	return authToken, refreshToken, nil
}

// DecodeToken func
func DecodeToken(token string, claims *Claims, jwtKey string) error {

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		return err
	}
	if !tkn.Valid {
		return jwt.ErrSignatureInvalid
	}
	return nil
}
