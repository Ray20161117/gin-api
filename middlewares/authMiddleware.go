/**
 * 鉴权中间件
 */
package middlewares

import (
	"gin-api/common/response"
	"gin-api/config/constant"
	"gin-api/pkg/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			response.Failed(c, int(response.ApiCode.NOAUTH),
				response.ApiCode.GetMessage(response.ApiCode.NOAUTH))
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.Failed(c, int(response.ApiCode.AUTHFORMATERROR),
				response.ApiCode.GetMessage(response.ApiCode.AUTHFORMATERROR))
			c.Abort()
			return
		}
		// todo 检验token
		mc, err := jwt.ValidateToken(parts[1])
		if err != nil {
			response.Failed(c, int(response.ApiCode.INVALIDTOKEN), response.ApiCode.GetMessage(response.ApiCode.INVALIDTOKEN))
			c.Abort()
			return
		}
		// 存用户信息
		c.Set(constant.ContextKeyUserObj, mc)
		c.Next()
	}
}
