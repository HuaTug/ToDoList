package middleware

import (
	"ToDoList/pkg/utils"
	"time"

	"github.com/gin-gonic/gin"
)

// JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := 200
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404
		} else {
			claims, err := utils.ParseToken(token)
			if err != nil {
				code = 403 //无权限，token是无权限的，是假的
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = 401 //Token无效
			}
		}
		if code != 200 {
			c.JSON(400, gin.H{
				"status": code,
				"msg":    "Token解析错误",
				//"data":   "可能是身份过期了，请重新登录",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
