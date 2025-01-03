/**
 * 部门实体类
 */
package entity

import "gin-api/common/utils"

type SysDept struct {
	Id         int         `gorm:"column:id;primary_key;auto_increment;comment:主键;NOT NULL" json:"id"` // 部门ID
	ParentId   int         `gorm:"column:parent_id;comment:父部门ID;NOT NULL" json:"parentId"`            // 父部门ID
	DeptType   int         `gorm:"column:dept_type;comment:部门类型;NOT NULL" json:"deptType"`             // 部门类型(1公司 2部门 3小组)
	DeptName   string      `gorm:"column:dept_name;comment:部门名称;NOT NULL" json:"deptName"`             // 部门名称
	DeptStatus int         `gorm:"column:dept_status;comment:部门状态;NOT NULL" json:"deptStatus"`         // 部门状态(1正常 2停用)
	CreateTime utils.HTime `gorm:"column:create_time;comment:创建时间;NOT NULL" json:"createTime"`         // 创建时间
	Children   []SysDept   `gorm:"-" json:"children"`                                                  // 子部门
}

func (SysDept) TableName() string {
	return "sys_dept"
}
