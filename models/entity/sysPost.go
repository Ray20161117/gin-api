/**
 * 岗位模型
 */
package entity

import "gin-api/common/utils"

// SysPost 岗位模型
type SysPost struct {
	Id         uint        `gorm:"column:id;primary_key;auto_increment;NOT NULL;comment:'岗位ID'" json:"id"`                   // 岗位ID
	PostCode   string      `gorm:"column:post_code;varchar(64);NOT NULL;comment:'岗位编码'" json:"postCode"`                     // 岗位编码
	PostName   string      `gorm:"column:post_name;varchar(50);NOT NULL;comment:'岗位名称'" json:"postName"`                     // 岗位名称
	PostStatus int         `gorm:"column:post_status;default:1;int(4);NOT NULL;comment:'岗位状态(1:正常 2:停用)'" json:"postStatus"` // 岗位状态(1:正常 2:停用)
	CreateTime utils.HTime `gorm:"column:create_time;NOT NULL;comment:'创建时间'" json:"createTime"`                             // 创建时间
	Remark     string      `gorm:"column:remark;varchar(500);comment:'备注'" json:"remark"`                                    // 备注
}

// TableName 表名
func (SysPost) TableName() string {
	return "sys_post"
}

// 批量伤处岗位
type BatchDelSysPostDto struct {
	Ids []uint `json:"ids"` // 岗位ID数组
}

// 修改岗位状态
type UpdateSysPostStatusDto struct {
	Id         uint `json:"id"`         // 岗位ID
	PostStatus int  `json:"postStatus"` // 岗位状态(1:正常 2:停用)
}

// 新增岗位
// swagger:model AddSysPost
type AddSysPostDto struct {
	PostCode   string `json:"postCode" validator:"required"`   // 岗位编码
	PostName   string `json:"postName" validator:"required"`   // 岗位名称
	PostStatus int    `json:"postStatus" validator:"required"` // 岗位状态(1:正常 2:停用)
	Remark     string `json:"remark"`                          // 备注
}

// 修改岗位
type UpdateSysPostDto struct {
	Id         uint   `json:"id" validator:"required"`         // 岗位ID
	PostCode   string `json:"postCode" validator:"required"`   // 岗位编码
	PostName   string `json:"postName" validator:"required"`   // 岗位名称
	PostStatus int    `json:"postStatus" validator:"required"` // 岗位状态(1:正常 2:停用)
	Remark     string `json:"remark"`                          // 备注
}
