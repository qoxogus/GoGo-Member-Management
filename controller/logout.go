package controller

import "github.com/gin-gonic/gin"

// Logout - 로그아웃
func Logout(c *gin.Context) {
	c.SetCookie("access-token", "", -1, "/", "localhost:3000", false, true)
	c.JSON(200, gin.H{
		"status":      200,
		"message":     "로그아웃이 완료되었습니다.",
		"accessToken": "null",
	})
}
