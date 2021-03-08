package controller

import (
	"Gin-api-server/database"
	"Gin-api-server/library/jwt"
	"fmt"

	"github.com/gin-gonic/gin"
)

// SigninParam - 파라미터 형식 구조체
type SigninParam struct {
	ID   string `json:"id" form:"id" query:"id"`
	Pw   string `json:"pw" form:"pw" query:"pw"`
	Name string `json:"name" query:"name"`
}

// Signin - 로그인 메서드
func Signin(c *gin.Context) {
	u := new(SigninParam)

	if err := c.Bind(u); err != nil {
		return
	}

	fmt.Println(u.ID, u.Pw, u.Name)

	User := &database.User{}
	err := database.DB.Where("user_id = ? AND pw = ?", u.ID, u.Pw).Find(User).Error
	if err != nil {
		c.JSON(500, gin.H{
			"status":       500,
			"message":      "일치하는 회원이 없습니다.",
			"refreshToken": "null",
			"accessToken":  "null",
		})
		return
	}

	fmt.Println(User.UserID, User.Name)

	refreshToken, err := jwt.CreateRefreshToken(User.UserID, User.Name)
	if err != nil {
		c.JSON(500, gin.H{
			"status":       500,
			"message":      "refreshtoken 생성 중 에러",
			"refreshToken": "null",
			"accessToken":  "null",
		})
		return
	}

	accessToken, err := jwt.CreateAccessToken(User.UserID, User.Name, User.IsManager)
	if err != nil {
		c.JSON(500, gin.H{
			"status":       500,
			"message":      "accesstoken 생성 중 에러",
			"refreshToken": refreshToken,
			"accessToken":  "null",
		})
		return
	}
	c.JSON(200, gin.H{
		"status": 200,
		"message": "토큰 발급 완료",
		"refreshToken": refreshToken,
		"accessToken":  accessToken,
	})
}
