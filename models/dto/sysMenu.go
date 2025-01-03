/**
 * 菜单数据层
 */
package dto

import (
	"gin-api/models/entity"
	"gin-api/pkg/db"
	"log"
)

// 菜单树结构
func buildMenuTree(menuList []entity.LeftMenuVoDto) (menuTree []entity.LeftMenuVoDto, err error) {
	for _, item := range menuList {
		if item.ParentId == 0 {
			item.MenuSvoList = findChildMenu(menuList, item.Id)
			menuTree = append(menuTree, item)
		}
	}
	return menuTree, nil
}

// 查找子菜单
func findChildMenu(menuList []entity.LeftMenuVoDto, menuId uint) (childMenuList []entity.LeftMenuVoDto) {
	for _, item := range menuList {
		if item.ParentId == menuId {
			childMenuList = append(childMenuList, item)
		}
	}
	return childMenuList
}

// 当前登录用户左侧菜单列表
func QueryLeftMenuList(Id uint) (leftMenuVo []entity.LeftMenuVoDto, err error) {
	// 使用链式调用来构建查询语句
	err = db.Db.Table("sys_menu sm").
		Select("sm.id, sm.menu_name, sm.url, sm.icon, sm.parent_id").
		Joins("JOIN sys_role_menu srm ON srm.menu_id = sm.id").
		Joins("JOIN sys_role sr ON sr.id = srm.role_id AND sr.status = ?", 1).
		Joins("JOIN sys_admin_role sar ON sar.role_id = sr.id").
		Where("sar.admin_id = ?", Id).
		Where("sm.menu_status = ?", 2).
		Where("sm.menu_type IN (?)", []int{1, 2}).
		Order("sm.sort").
		Find(&leftMenuVo).Error

	// 检查并处理错误
	if err != nil {
		log.Printf("Error querying left menu list for admin ID %d: %v", Id, err)
		return nil, err
	}

	// 构建菜单树
	leftMenuVo, err = buildMenuTree(leftMenuVo)
	if err != nil {
		log.Printf("Error building menu tree for admin ID %d: %v", Id, err)
		return nil, err
	}
	return leftMenuVo, nil
}

// 当前登录用户左侧菜单级列表
func QueryMenuVoList(AdminId, MenuId uint) (menuSvo []entity.MenuSvoDto, err error) {
	const status, menuStatus, menuType = 1, 2, 2

	err = db.Db.Table("sys_menu sm").
		Select("sm.menu_name, sm.icon, sm.url").
		Joins("LEFT JOIN sys_role_menu srm ON sm.id = srm.menu_id").
		Joins("LEFT JOIN sys_role sr ON sr.id = srm.role_id").
		Joins("LEFT JOIN sys_admin_role sar ON sar.role_id = sr.id").
		Joins("LEFT JOIN sys_admin sa ON sa.id = sar.admin_id").
		Where("sr.status = ?", status).
		Where("sm.menu_status = ?", menuStatus).
		Where("sm.menu_type = ?", menuType).
		Where("sm.parent_id = ?", MenuId).
		Where("sa.id = ?", AdminId).
		Order("sm.sort").
		Scan(&menuSvo).Error

	// 检查并处理错误
	if err != nil {
		log.Printf("Error querying menu list for admin ID %d and menu ID %d: %v", AdminId, MenuId, err)
		return nil, err
	}

	return menuSvo, nil
}

// 当前登录用户的权限列表
func QueryPermissionValueList(Id uint) (valueVo []entity.ValueVoDto, err error) {
	const status, menuStatus, menuType uint = 1, 2, 1
	// 使用链式调用来构建查询语句，并检查错误
	err = db.Db.Table("sys_menu sm").
		Select("sm.value").
		Joins("LEFT JOIN sys_role_menu srm ON sm.id = srm.menu_id").
		Joins("LEFT JOIN sys_role sr ON sr.id = srm.role_id").
		Joins("LEFT JOIN sys_admin_role sar ON sar.role_id = sr.id").
		Joins("LEFT JOIN sys_admin sa ON sa.id = sar.admin_id").
		Where("sr.status = ?", status).
		Where("sm.menu_status = ?", menuStatus).
		Where("sm.menu_type != ?", menuType).
		Where("sa.id = ?", Id).
		Scan(&valueVo).Error

	// 检查并处理错误
	if err != nil {
		log.Printf("Error querying permission value list for admin ID %d: %v", Id, err)
		return nil, err
	}

	return valueVo, nil
}
