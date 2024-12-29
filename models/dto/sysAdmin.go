/**
 * 用户数据层
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

// 用户信息获取--登陆
func SysAdminDetail(dto entity.LoginDto) (sysAdmin entity.SysAdmin, err error) {
	result := db.Db.Where("username = ?", dto.Username).First(&sysAdmin)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return sysAdmin, errors.New("用户未找到")
		}
		return sysAdmin, errors.New("查询用户时发生错误: " + result.Error.Error())
	}
	return sysAdmin, nil
}

// 根据用户名查询用户
func GetSysAdminByUsername(username string) (sysAdmin entity.SysAdmin) {
	db.Db.Where("username =?", username).First(&sysAdmin)
	return sysAdmin
}

// 新增用户
func AddSysAdmin(dto entity.AddSysAdminDto) bool {
	// 开启事务
	tx := db.Db.Begin()
	if tx.Error != nil {
		return false
	}

	// 检查用户名是否已存在
	var existingSysAdmin entity.SysAdmin
	if err := tx.Where("username = ?", dto.Username).First(&existingSysAdmin).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 用户未找到，继续添加新用户
		} else {
			tx.Rollback()
			return false
		}
	} else {
		tx.Rollback()
		return false
	}

	// 创建SysAdmin记录
	sysAdmin := entity.SysAdmin{
		PostId:     dto.PostId,
		DeptId:     dto.DeptId,
		Username:   dto.Username,
		Password:   utils.EncryptionMd5(dto.Password),
		Phone:      dto.Phone,
		Email:      dto.Email,
		Note:       dto.Note,
		Status:     dto.Status,
		CreateTime: utils.HTime{Time: time.Now()},
	}
	if err := tx.Create(&sysAdmin).Error; err != nil {
		tx.Rollback()
		return false
	}

	// 创建SysAdminRole记录
	sysAdminRole := entity.SysAdminRole{
		AdminId: sysAdmin.ID,
		RoleId:  dto.RoleId,
	}
	if err := tx.Create(&sysAdminRole).Error; err != nil {
		tx.Rollback()
		return false
	}

	tx.Commit() // 提交事务
	return true
}

// 根据id删除用户
func DeleteSysAdminById(dto entity.SysAdminIdDto) {
	db.Db.First(&entity.SysAdmin{}, dto.Id)
	db.Db.Delete(&entity.SysAdmin{}, dto.Id)
	db.Db.Where("admin_id = ?", dto.Id).Delete(&entity.SysAdminRole{})
}

// 修改用户状态
func UpdateSysAdminStatus(dto entity.UpdateSysAdminStatusDto) {
	var sysAdmin entity.SysAdmin
	db.Db.First(&sysAdmin, dto.Id)
	sysAdmin.Status = dto.Status
	db.Db.Save(&sysAdmin)
}

// 重置密码
func ResetSysAdminPassword(dto entity.ResetSysAdminPasswordDto) {
	var sysAdmin entity.SysAdmin
	db.Db.First(&sysAdmin, dto.Id)
	sysAdmin.Password = utils.EncryptionMd5(dto.Password)
	db.Db.Save(&sysAdmin)
}

// 分页查询用户列表
func GetSysAdminList(PageSize, PageNum int, Username, Status, BeginTime, EndTime string) (sysAdminVo []entity.SysAdminVo, count int64) {
	curDb := db.Db.Table("sys_admin").
		Select("sys_admin.*, sys_post.post_name, sys_role.role_name, sys_dept.dept_name").
		Joins("LEFT JOIN sys_post ON sys_admin.post_id = sys_post.id").
		Joins("LEFT JOIN sys_admin_role ON sys_admin.id = sys_admin_role.admin_id").
		Joins("LEFT JOIN sys_role ON sys_role.id = sys_admin_role.role_id").
		Joins("LEFT JOIN sys_dept ON sys_dept.id = sys_admin.dept_id")
	if Username != "" {
		curDb = curDb.Where("sys_admin.username = ?", Username)
	}
	if Status != "" {
		curDb = curDb.Where("sys_admin.status = ?", Status)
	}
	if BeginTime != "" && EndTime != "" {
		curDb = curDb.Where("sys_admin.create_time BETWEEN ? AND ?", BeginTime, EndTime)
	}
	curDb.Count(&count)
	curDb.Limit(PageSize).Offset((PageNum - 1) * PageSize).Order("sys_admin.create_time DESC").Find(&sysAdminVo)
	return sysAdminVo, count
}

// 修改个人信息
func UpdatePersonal(dto entity.UpdatePersonalDto) (sysAdmin entity.SysAdmin) {
	db.Db.First(&sysAdmin, dto.Id)
	if dto.Icon != "" {
		sysAdmin.Icon = dto.Icon
	}
	if dto.Username != "" {
		sysAdmin.Username = dto.Username
	}
	if dto.Nickname != "" {
		sysAdmin.Nickname = dto.Nickname
	}
	if dto.Phone != "" {
		sysAdmin.Phone = dto.Phone
	}
	if dto.Email != "" {
		sysAdmin.Email = dto.Email
	}
	db.Db.Save(&sysAdmin)
	return sysAdmin
}

// 修改个人密码
func UpdatePersonalPassword(dto entity.UpdatePersonalPasswordDto) (sysAdmin entity.SysAdmin) {
	db.Db.First(&sysAdmin, dto.Id)
	sysAdmin.Password = dto.NewPassword
	db.Db.Save(&sysAdmin)
	return sysAdmin
}
