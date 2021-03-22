package controller

import "github.com/gin-gonic/gin"

// Logout - 로그아웃
func Logout(c *gin.Context) {
	// c.SetCookie("access-token", "", -1, "/", "localhost:3000", false, true)

	//클라에서 토큰지워주는 방법 사용

	c.JSON(200, gin.H{
		"status":       200,
		"message":      "로그아웃",
		"accessToken":  "null",
		"refreshToken": "null",
	})
	return
}
