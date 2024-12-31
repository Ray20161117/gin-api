/**
 * 岗位的控制层
 */
package controllers

import (
	"gin-api/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

//var sysPost entity.SysPost

// 岗位列表
// @Summary 岗位列表(分页)
// @Tags 岗位管理
// @Produce  json
// @Description 获取岗位列表
// @Param page query int false "页码"
// @Param pageSize query int false "每页数量"
// @Param postName query string false "岗位名称"
// @Param postStatus query string false "岗位状态:1启用,2禁用"
// @Param beginTime query string false "开始时间"
// @Param endTime query string false "结束时间"
// @Success 200 {object} response.Result
// @Router /api/post/list [get]
// @Security ApiKeyAuth
func GetSysPostList(c *gin.Context) {
	Page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	PostName := c.Query("postName")
	PostStatus := c.Query("postStatus")
	BeginTime := c.Query("beginTime")
	EndTime := c.Query("endTime")
	services.SysPostService().GetList(c, Page, pageSize, PostName, PostStatus, BeginTime, EndTime)
}
