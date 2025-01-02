/**
 * 岗位的控制层
 */
package controllers

import (
	"gin-api/models/entity"
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

// 新增岗位
// @Tags 岗位管理
// @Summary 新增岗位
// @Produce  json
// @Description 新增岗位
// @Param data body entity.AddSysPostDto true "请求参数"
// @Success 200 {object} response.Result
// @Router /api/post/add [post]
// @Security ApiKeyAuth
func AddSysPost(c *gin.Context) {
	var dto entity.AddSysPostDto
	_ = c.ShouldBindJSON(&dto)
	services.SysPostService().AddSysPost(c, dto)
}

// 编辑岗位
// @Tags 岗位管理
// @Summary 编辑岗位
// @Produce  json
// @Description 编辑岗位
// @Param data body entity.UpdateSysPostDto true "请求参数"
// @Success 200 {object} response.Result
// @Router /api/post/update [put]
// @Security ApiKeyAuth
func UpdateSysPost(c *gin.Context) {
	var dto entity.UpdateSysPostDto
	_ = c.ShouldBindJSON(&dto)
	services.SysPostService().UpdateSysPost(c, dto)
}

// 获取岗位详情
// @Tags 岗位管理
// @Summary 获取岗位详情
// @Produce  json
// @Description 获取岗位详情
// @Param postId path int true "岗位ID"
// @Success 200 {object} response.Result
// @Router /api/post/info/{postId} [get]
// @Security ApiKeyAuth
func GetSysPostDetail(c *gin.Context) {
	postId, _ := strconv.Atoi(c.Param("postId"))
	services.SysPostService().GetSysPostDetail(c, postId)
}

// 删除岗位
// @Tags 岗位管理
// @Summary 删除岗位
// @Produce  json
// @Description 删除岗位
// @Param postId path int true "岗位ID"
// @Success 200 {object} response.Result
// @Router /api/post/del/{postId} [delete]
// @Security ApiKeyAuth
func DelSysPostById(c *gin.Context) {
	postId, _ := strconv.Atoi(c.Param("postId"))
	services.SysPostService().DeleteSysPostById(c, postId)
}

// 批量删除岗位
// @Tags 岗位管理
// @Summary 批量删除岗位
// @Produce  json
// @Description 批量删除岗位
// @Param data body entity.BatchDelSysPostDto true "请求参数" example:"{ids: [1,2,3]}"
// @Success 200 {object} response.Result
// @Router /api/post/batchDel [delete]
// @Security ApiKeyAuth
func BatchDelSysPostByIds(c *gin.Context) {
	var dto entity.BatchDelSysPostDto
	_ = c.ShouldBindJSON(&dto)
	services.SysPostService().DeleteSysPostByIds(c, dto)
}
