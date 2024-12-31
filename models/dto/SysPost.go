/**
 * SysPosts数据层
 */
package dto

import (
	"gin-api/models/entity"
	"gin-api/pkg/db"
)

// 分页查询岗位列表
func GetSysPostList(PageNum, PageSize int, PostName, PostStatus, BeginTime, EndTime string) (sysPost []entity.SysPost, count int64, err error) {
	curDB := db.Db.Table("sys_post")

	if PostName != "" {
		curDB = curDB.Where("post_name like ?", "%"+PostName+"%")
	}
	if BeginTime != "" && EndTime != "" {
		curDB = curDB.Where("create_time BETWEEN ? AND ?", BeginTime, EndTime)
	}
	if PostStatus != "" {
		curDB = curDB.Where("post_status = ?", PostStatus)
	}

	// 使用同一个 curDB 进行 Count 和 Find 操作
	if err = curDB.Count(&count).Error; err != nil {
		return
	}
	if err = curDB.Offset((PageNum - 1) * PageSize).Limit(PageSize).Find(&sysPost).Error; err != nil {
		return
	}

	return sysPost, count, nil
}
