package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jwtMethod interface {
	CreateRefreshToken()
	CreateAccessToken()
	VerifyRefreshToken()
	VerifyAccessToken()
}

// CreateRefreshToken : Middleware that create RefreshToken
func CreateRefreshToken(Id, Name string) (string, error) {
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	claims := refreshToken.Claims.(jwt.MapClaims)
	claims["ID"] = Id
	claims["Name"] = Name
	claims["exp"] = time.Now().Add(time.Hour * 720).Unix()

	t, err := refreshToken.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return t, nil
}

// CreateAccessToken : Middleware that create AccessToken
func CreateAccessToken(Id, Name string, IsManager bool) (string, error) {
	accessToken := jwt.New(jwt.SigningMethodHS256)
	claims := accessToken.Claims.(jwt.MapClaims)
	claims["ID"] = Id
	claims["Name"] = Name
	claims["IsManager"] = IsManager
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	t, err := accessToken.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return t, nil
}
