/**
 * 操作日志中间件
 */
package middlewares

import (
	"gin-api/common/utils"
	"gin-api/models/dto"
	"gin-api/models/entity"
	"gin-api/pkg/jwt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := strings.ToLower(c.Request.Method)
		sysAdmin, _ := jwt.GetAdmin(c)
		if method != "get" {
			log := entity.SysOperationLog{
				AdminId:    sysAdmin.ID,
				Username:   sysAdmin.Username,
				Method:     method,
				Ip:         c.ClientIP(),
				Url:        c.Request.URL.Path,
				CreateTime: utils.HTime{Time: time.Now()},
			}
			dto.CreateSysOperationLog(log)
		}
	}
}
