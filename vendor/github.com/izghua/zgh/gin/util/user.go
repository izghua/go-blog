package ginutil

import (
	"github.com/gin-gonic/gin"
)

// 获取用户ID
func GetUserID(c *gin.Context) int {
	userID, exists := c.Get("userId")
	if !exists {
		return 0
	}
	if value, ok := userID.(int); ok {
		return value
	}
	return 0
}