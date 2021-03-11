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
	// c.Get("user").(*jwt.Token)
	ctoken, err := c.Request.Cookie("access-token")
	if err != nil {
		c.JSON(401, gin.H{
			"status":  401,
			"message": "Authentication failed",
		})
		c.Abort()
		return
	}
	tknstr := ctoken.Value

	fmt.Println(ctoken)                     //쿠키에서 받아온 값
	fmt.Println("token string : " + tknstr) // 쿠키에서 value로 추출해온 값

	if tknstr == "" {
		c.JSON(401, gin.H{
			"status":  401,
			"message": "token is None.",
		})
		c.Abort()
		return
	}

	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tknstr, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		c.JSON(401, gin.H{
			"status":  401,
			"message": "토큰 인증 실패.",
		})
	}

	fmt.Printf("token : %v\n", token)

	for key, val := range claims {
		fmt.Printf("Key : %v, value : %v\n", key, val)
	}
	c.JSON(200, gin.H{
		"status":  200,
		"message": "토큰 인증 완료.",
	})
	return
}
