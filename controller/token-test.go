package controller

import "github.com/gin-gonic/gin"

//TokenTest - 토큰 검증 테스트용 코드
func TokenTest(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hi~ token test",
	})
}
