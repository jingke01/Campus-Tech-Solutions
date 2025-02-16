package middleware

import (
	"github.com/gin-gonic/gin"
	"temp/tools"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			tools.UnauthorizationRequest(c, "请登录")
			c.Abort()
			return
		}

		claims, err := tools.ParseToken(tokenString)

		if err != nil {
			tools.UnauthorizationRequest(c, "登录失败")
			c.Abort()
			return
		}

		c.Set("user_id", claims["user_id"])
	}
}
