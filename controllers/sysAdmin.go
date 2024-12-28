/**
 * 用户控制层
 */
package controllers

import (
	"gin-api/models/entity"
	"gin-api/services"

	"github.com/gin-gonic/gin"
)

// @Summary 用户登录接口
// @Tags 用户管理
// @Produce json
// @Description 用户登录接口
// @Param data body entity.LoginDto true "data"
// @Success 200 {object} response.Result
// @router /api/login [post]
func Login(c *gin.Context) {
	var dto entity.LoginDto
	_ = c.BindJSON(&dto)
	services.SysAdminService().Login(c, dto)
}
