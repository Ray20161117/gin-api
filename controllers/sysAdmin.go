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

// 新增用户
// @Tags 用户管理
// @Summary 新增用户接口
// @Produce json
// @Description 新增用户接口
// @Param data body entity.AddSysAdminDto true "data"
// @Success 200 {object} response.Result
// @router /api/admin/add [post]
// @Security ApiKeyAuth
func CreateSysAdmin(c *gin.Context) {
	var dto entity.AddSysAdminDto
	_ = c.BindJSON(&dto)
	services.SysAdminService().AddSysAdmin(c, dto)
}
