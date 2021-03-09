package jwt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type jwtMethod interface {
	CreateRefreshToken()
	CreateAccessToken()
	VerifyRefreshToken()
	VerifyAccessToken()
}

// CreateRefreshToken - Middleware that create RefreshToken
func CreateRefreshToken(Name string) (string, error) {
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	claims := refreshToken.Claims.(jwt.MapClaims)
	claims["Name"] = Name
	claims["exp"] = time.Now().Add(time.Hour * 720).Unix()

	t, err := refreshToken.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return t, nil
}

// CreateAccessToken - Middleware that create AccessToken
func CreateAccessToken(Name string, IsManager bool) (string, error) {
	accessToken := jwt.New(jwt.SigningMethodHS256)
	claims := accessToken.Claims.(jwt.MapClaims)
	claims["Name"] = Name
	claims["IsManager"] = IsManager
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	t, err := accessToken.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return t, nil
}

// VerifyRefreshToken - Middleware that verify RefreshToken
func VerifyRefreshToken() {

}

// VerifyAccessToken - Middleware that verify AccessToken
func VerifyAccessToken(c *gin.Context) {
	token, err := c.Request.Cookie("access-token")
	if err != nil {
		c.JSON(401, gin.H{
			"status":  401,
			"message": "Authentication failed",
		})
		c.Abort()
		return
	}
	tknstr := token.Value

	fmt.Println(token)
	fmt.Println("token string : " + tknstr)

	if tknstr == "" {
		c.JSON(401, gin.H{
			"status":  401,
			"message": "token is None.",
		})
		c.Abort()
		return
	}

	// claims := tknstr.Claims(jwt.MapClaims)
	// Name := claims["Name"].(string)
	// IsManager := claims["IsManager"].(bool)
	// c.Set("Name", Name)
	// c.Set("IsManager", IsManager)
}
