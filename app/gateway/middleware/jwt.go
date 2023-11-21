package middleware

import (
	"time"

	"github.com/gin-gonic/gin"

	"one_ser.com/errno"
	"one_ser.com/jwt"
)

// JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		code = 200
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404
			c.JSON(200, gin.H{
				"status": code,
				"msg":    errno.GetMsg(code),
				"data":   data,
			})
			c.Abort()
		}
		claims, err := jwt.ParseToken(token)
		if err != nil {
			code = errno.ErrorAuthCheckTokenFail
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = errno.ErrorAuthCheckTokenTimeout
		}
		if code != errno.SUCCESS {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    errno.GetMsg(code),
				"data":   data,
			})
			c.Abort()
			return
		}
		// c.Request = c.Request.WithContext(ctl.NewContext(c.Request.Context(), &ctl.UserInfo{Id: claims.UserID}))
		// ctl.InitUserInfo(c.Request.Context())
		c.Next()
	}
}
