/**
 * SysPosts数据层
 */
package dto

import (
	"errors"
	"gin-api/common/utils"
	"gin-api/models/entity"
	"gin-api/pkg/db"
	"time"

	"gorm.io/gorm"
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

func AddSysPost(addSysPostDto entity.AddSysPostDto) bool {
	tx := db.Db.Begin()
	if tx.Error != nil {
		return false
	}
	var exittingSysPost entity.SysPost
	if err := tx.Where("post_code = ?", addSysPostDto.PostCode).Or("post_name = ?", addSysPostDto.PostName).First(&exittingSysPost).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			addSysPost := entity.SysPost{
				PostCode:   addSysPostDto.PostCode,
				PostName:   addSysPostDto.PostName,
				PostStatus: addSysPostDto.PostStatus,
				CreateTime: utils.HTime{Time: time.Now()},
				Remark:     addSysPostDto.Remark,
			}
			if err := tx.Create(&addSysPost).Error; err != nil {
				tx.Rollback()
				return false
			} else {
				tx.Commit() // 提交事务
				return true
			}
		} else {
			tx.Rollback()
			return false
		}
	} else {
		tx.Rollback()
		return false
	}
}