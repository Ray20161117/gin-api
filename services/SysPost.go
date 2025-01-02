/**
 * SysPost业务层
 */
package services

import (
	"gin-api/common/response"
	"gin-api/common/utils"
	"gin-api/models/dto"
	"gin-api/models/entity"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// ISysPostService 岗位业务接口
type ISysPostService interface {
	GetList(c *gin.Context, Page, PageSize int, PostName, PostStatus, BeginTime, EndTime string)
	AddSysPost(c *gin.Context, addSysPostDto entity.AddSysPostDto)
	UpdateSysPost(c *gin.Context, updateSysPostDto entity.UpdateSysPostDto)
	GetSysPostDetail(c *gin.Context, id int)
	DeleteSysPostById(c *gin.Context, id int)
}

// SysPostServiceImpl 实现接口业务的结构体
type SysPostServiceImpl struct{}

// 实例化实现接口的结构体
var sysPostService = SysPostServiceImpl{}

// SysPostService工厂方法,返回接口的一个指针，指向SysPostService实例,实现业务正常被其它模块调用，同时隐藏了具体的实现细节
func SysPostService() ISysPostService {
	return &sysPostService
}

// 以下为具体的实现方法
// 分页查询岗位列表
func (s SysPostServiceImpl) GetList(c *gin.Context, PageNum, PageSize int,
	PostName, PostStatus, BeginTime, EndTime string) {
	if PageSize < 1 {
		PageSize = 10
	}
	if PageNum < 1 {
		PageNum = 1
	}
	sysPost, count, err := dto.GetSysPostList(PageNum, PageSize, PostName, PostStatus,
		BeginTime, EndTime)
	if err != nil {
		// 错误处理逻辑，例如返回错误信息给客户端
		response.Failed(c, int(response.ApiCode.FAILED), response.ApiCode.GetMessage(response.ApiCode.FAILED))
		return
	}
	response.Success(c, map[string]interface{}{"total": count, "pageSize": PageSize,
		"pageNum": PageNum, "list": sysPost})
}

// 新增岗位
func (s SysPostServiceImpl) AddSysPost(c *gin.Context, addSysPostDto entity.AddSysPostDto) {
	if err := validator.New().Struct(addSysPostDto); err != nil {
		if firstError := err.(validator.ValidationErrors)[0]; firstError != nil {
			msg := utils.TranslateError(firstError.Field(), firstError.Tag(), firstError.Param())
			if msg != "" {
				response.Failed(c, int(response.ApiCode.INVALID_PARAMS), msg)
				return
			}
		}
		response.Failed(c, int(response.ApiCode.INVALID_PARAMS), response.ApiCode.GetMessage(response.ApiCode.INVALID_PARAMS))
		return
	}
	bool := dto.AddSysPost(addSysPostDto)
	if !bool {
		response.Failed(c, int(response.ApiCode.FAILED), response.ApiCode.GetMessage(response.ApiCode.FAILED))
	}
	response.Success(c, nil)
}

// 编辑岗位
func (s SysPostServiceImpl) UpdateSysPost(c *gin.Context, updateSysPostDto entity.UpdateSysPostDto) {
	if err := validator.New().Struct(updateSysPostDto); err != nil {
		if firstError := err.(validator.ValidationErrors)[0]; firstError != nil {
			msg := utils.TranslateError(firstError.Field(), firstError.Tag(), firstError.Param())
			if msg != "" {
				response.Failed(c, int(response.ApiCode.INVALID_PARAMS), msg)
				return
			}
		}
		response.Failed(c, int(response.ApiCode.INVALID_PARAMS), response.ApiCode.GetMessage(response.ApiCode.INVALID_PARAMS))
		return
	}
	bool := dto.UpdateSysPost(updateSysPostDto)
	if !bool {
		response.Failed(c, int(response.ApiCode.FAILED), response.ApiCode.GetMessage(response.ApiCode.FAILED))
	}
	response.Success(c, nil)
}

// 岗位详情
func (s SysPostServiceImpl) GetSysPostDetail(c *gin.Context, id int) {
	sysPost, err := dto.GetSysPostDetail(id)
	if err != nil {
		response.Failed(c, int(response.ApiCode.FAILED), response.ApiCode.GetMessage(response.ApiCode.FAILED))
		return
	}
	response.Success(c, sysPost)
}

// 删除岗位(单个)
func (s SysPostServiceImpl) DeleteSysPostById(c *gin.Context, id int) {
	bool := dto.DeleteSysPostById(id)
	if bool {
		response.Success(c, nil)
	} else {
		response.Failed(c, int(response.ApiCode.FAILED), response.ApiCode.GetMessage(response.ApiCode.FAILED))
		return
	}
}
