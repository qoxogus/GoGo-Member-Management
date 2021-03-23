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
	CreateReissuanceToken()
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
func VerifyRefreshToken(c *gin.Context) {
	//refreshToken을 DB에 넣는 방법도있다. (조금 더 안전하다.) accessToken을 보내며 재발급 요청 -> DB에서 refreshToken을 가져와 검증. -> 검증완료시 accessToken재발급.
	htoken := c.GetHeader("user-refresh-token") //리프레쉬토큰

	fmt.Println(htoken)

	if htoken == "" {
		c.JSON(401, gin.H{
			"status":  401,
			"message": "token is None.",
		})
		return
	}

	fmt.Println(htoken)

	if htoken == "" {
		c.JSON(401, gin.H{
			"status":  401,
			"message": "refresh token is None. (다시 로그인하세요.)",
		})
		return
	}

	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(htoken, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		c.JSON(401, gin.H{
			"status":  401,
			"message": "refreshToken이 만료되었습니다. 다시 로그인하세요.",
		})
		return
	}

	fmt.Printf("token : %v\n", token)

	for key, val := range claims {
		fmt.Printf("Key : %v, value : %v\n", key, val)
	}

	c.JSON(200, gin.H{
		"status":  200,
		"message": "refreshToken 검증 완료.",
	})
	return
}

// VerifyAccessToken - Middleware that verify AccessToken
func VerifyAccessToken(c *gin.Context) {
	htoken := c.GetHeader("user-token") //엑세스토큰

	fmt.Println(htoken)

	if htoken == "" {
		c.JSON(401, gin.H{
			"status":  401,
			"message": "token is None.",
		})
		return
	}

	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(htoken, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		c.JSON(401, gin.H{
			"status":  401,
			"message": "토큰 인증 실패. 토큰을 재발급 받으세요.(한번 재발급 받았다면 다시 로그인 그렇지않다면 재발급요청)", // Client에서 이 메세지를 받고 accessToken을 재발급 받기위해 refreshToken을 보내며 재발급 요청.
		})
		return
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

func CreateReissuanceToken(c *gin.Context) {
	htoken := c.GetHeader("user-token") //토큰

	fmt.Println(htoken)

	if htoken == "" {
		c.JSON(401, gin.H{
			"status":  401,
			"message": "token is None.",
		})
		return
	}

	claims := jwt.MapClaims{}

	token, _ := jwt.ParseWithClaims(htoken, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	fmt.Println(token)

	for key, val := range claims {
		fmt.Printf("Key : %v, value : %v\n", key, val)
	}

	Name := claims["Name"].(string)
	IsManager := claims["IsManager"].(bool)
	accessToken, err := CreateAccessToken(Name, IsManager)

	if err != nil {
		c.JSON(500, gin.H{
			"status":  500,
			"message": "accessToken 생성중 에러",
		})
		return
	}
	c.JSON(200, gin.H{
		"status":      200,
		"message":     "accessToken 재발급 완료.",
		"accessToken": accessToken,
	})
	return
}
