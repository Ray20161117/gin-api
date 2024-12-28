/**
 * 验证码接口
 */
package controllers

import (
	"gin-api/common/response"
	"gin-api/services"

	"github.com/gin-gonic/gin"
)

// 验证码
// @Summary 验证码接口
// @Tags 验证码
// @Produce json
// @Description 验证码接口
// @Success 200 {object} response.Result
// @Router /api/captcha [get]
func Captcha(c *gin.Context) {
	id, base64Image := services.CaptMake()
	response.Success(c, map[string]interface{}{"idKey": id, "image": base64Image})
}
