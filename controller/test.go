package controller

import "github.com/gin-gonic/gin"

//ServerTest - 서버 테스트용 코드
func ServerTest(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "welcome",
	})
}
