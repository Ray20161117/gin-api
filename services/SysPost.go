/**
 * SysPost业务层
 */
package services

import (
	"gin-api/common/response"
	"gin-api/models/dto"

	"github.com/gin-gonic/gin"
)

// ISysPostService 岗位业务接口
type ISysPostService interface {
	GetList(c *gin.Context, Page, PageSize int, PostName, PostStatus, BeginTime, EndTime string)
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
