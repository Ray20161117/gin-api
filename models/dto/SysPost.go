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

// 新增岗位
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

// 更新岗位
func UpdateSysPost(updateSysPostDto entity.UpdateSysPostDto) bool {
	tx := db.Db.Begin()
	if tx.Error != nil {
		return false
	}
	var exittingSysPost entity.SysPost
	if err := tx.Where("post_code = ?", updateSysPostDto.PostCode).
		Or("post_name = ?", updateSysPostDto.PostName).
		Not("id = ?", updateSysPostDto.Id).
		First(&exittingSysPost).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			var updateSysPost entity.SysPost
			if err := tx.First(&updateSysPost, updateSysPostDto.Id).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return false
				}
				tx.Rollback()
				return false
			}
			updateSysPost.PostCode = updateSysPostDto.PostCode
			updateSysPost.PostName = updateSysPostDto.PostName
			updateSysPost.PostStatus = updateSysPostDto.PostStatus
			updateSysPost.Remark = updateSysPostDto.Remark

			if err := tx.Save(&updateSysPost).Error; err != nil {
				tx.Rollback()
				return false
			}
			tx.Commit() // 提交事务
			return true
		} else {
			tx.Rollback()
			return false
		}
	} else {
		tx.Rollback()
		return false
	}
}

// 岗位详情
func GetSysPostDetail(id int) (sysPost entity.SysPost, err error) {
	if err := db.Db.First(&sysPost, id).Error; err != nil {
		return sysPost, err
	}
	return sysPost, nil
}

// 删除岗位
func DeleteSysPostById(id int) bool {
	tx := db.Db.Begin()
	if tx.Error != nil {
		return false
	}
	if err := tx.Delete(&entity.SysPost{}, id).Error; err != nil {
		tx.Rollback()
		return false
	}
	tx.Commit() // 提交事务
	return true
}

// 批量删除岗位
func DeleteSysPostByIds(ids []uint) bool {
	if len(ids) == 0 {
		return false // 如果没有提供任何ID，则无需进行删除操作，直接返回nil
	}
	tx := db.Db.Begin()
	if tx.Error != nil {
		return false
	}
	if err := tx.Delete(&entity.SysPost{}, ids).Error; err != nil {
		tx.Rollback()
		return false
	}
	tx.Commit() // 提交事务
	return true
}

// 变更岗位状态
func ChangeSysPostStatus(id int, status int) bool {
	tx := db.Db.Begin()
	if tx.Error != nil {
		return false
	}
	var sysPost entity.SysPost
	if err := tx.First(&sysPost, id).Error; err != nil {
		tx.Rollback()
		return false
	}
	sysPost.PostStatus = status
	if err := tx.Save(&sysPost).Error; err != nil {
		tx.Rollback()
		return false
	}
	tx.Commit() // 提交事务
	return true
}
