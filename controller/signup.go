package controller

import (
	"Gin-api-server/database"
	"fmt"

	"github.com/gin-gonic/gin"
)

// SignUpParam - 파라미터 형식 구조체
type SignUpParam struct {
	ID        string `json:"id" form:"id" query:"id"`
	Pw        string `json:"pw" form:"pw" query:"pw"`
	Name      string `json:"name" form:"name" query:"name"`
	IsManager bool   `json:"manager" form:"manager" query:"manager"`
}

// SignUp - 회원가입
func SignUp(c *gin.Context) {
	u := new(SignUpParam)

	if err := c.Bind(u); err != nil {
		return
	}

	fmt.Println(u.ID, u.Pw, u.Name, u.IsManager)

	if u.ID == "" || u.Pw == "" || u.Name == "" {
		c.JSON(400, gin.H{
			"status":  400,
			"message": "모든 값을 입력해주세요.",
		})
		return
	}
	User := &database.User{}
	err := database.DB.Where("user_id = ?", u.ID).Find(User).Error
	if err == nil {
		c.JSON(400, gin.H{
			"status":  400,
			"message": "이미 사용중인 아이디입니다.",
		})
		return
	}

	User = &database.User{UserID: u.ID, Pw: u.Pw, Name: u.Name, IsManager: u.IsManager}
	err = database.DB.Create(User).Error
	if err != nil {
		c.JSON(500, gin.H{
			"status":  500,
			"message": "회원가입 실패",
		})
		return
	}
	c.JSON(200, gin.H{
		"status":  200,
		"message": "회원가입이 완료되었습니다.",
	})
	return
}
