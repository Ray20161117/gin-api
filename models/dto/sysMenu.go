/**
 * 菜单数据层
 */
package dto

import (
	"errors"
	"gin-api/models/entity"
	"gin-api/pkg/db"
	"log"
)

// 菜单树结构
func buildMenuTree(menuList []entity.SysMenu) (menuTree []entity.LeftMenuVoDto, err error) {
	// 使用map来存储以ParentId为键的子菜单列表
	childrenMap := make(map[uint][]entity.SysMenu)
	// 遍历menuList，填充childrenMap，并创建topMenusDto
	for _, item := range menuList {
		if item.ParentId == 0 {
			// 创建一个新的LeftMenuVoDto实例，并复制需要的字段
			topMenuDto := entity.LeftMenuVoDto{
				Id:          item.ID,
				MenuName:    item.MenuName,
				Icon:        item.Icon,
				Url:         item.Url,
				MenuSvoList: []entity.MenuSvoDto{},
				// 复制其他的所需字段...
			}
			menuTree = append(menuTree, topMenuDto)
		}
		childrenMap[item.ParentId] = append(childrenMap[item.ParentId], item)
	}

	// 如果没有找到顶层菜单，返回错误
	if len(menuTree) == 0 {
		return nil, errors.New("no top-level menus found")
	}

	// 遍历顶层菜单，并查找其子菜单
	for i, item := range menuTree {
		subMenu, err := buildSubMenuTree(item.Id, childrenMap)
		if err != nil {
			return nil, err
		}
		// 将子菜单赋值给当前顶层菜单
		menuTree[i].MenuSvoList = subMenu
	}

	return menuTree, nil
}

// 递归构建子菜单树
func buildSubMenuTree(menuId uint, childrenMap map[uint][]entity.SysMenu) (subMenuTree []entity.MenuSvoDto, err error) {
	childMenus, exists := childrenMap[menuId]
	if !exists {
		return nil, nil // 如果没有子菜单，返回空列表而不是错误
	}

	for _, item := range childMenus {
		// 创建一个新的LeftMenuVoDto实例，并复制需要的字段
		childMenuDto := entity.MenuSvoDto{
			MenuName: item.MenuName,
			Icon:     item.Icon,
			Url:      item.Url,
			// 复制其他的所需字段...
		}
		subMenuTree = append(subMenuTree, childMenuDto)
	}

	return subMenuTree, nil
}

// 当前登录用户左侧菜单列表
func QueryLeftMenuList(Id uint) (leftMenuVo []entity.LeftMenuVoDto, err error) {
	var sysMenu []entity.SysMenu
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
		Find(&sysMenu).Error

	// 检查并处理错误
	if err != nil {
		log.Printf("Error querying left menu list for admin ID %d: %v", Id, err)
		return nil, err
	}

	// 构建菜单树
	leftMenuVo, err = buildMenuTree(sysMenu)
	if err != nil {
		log.Printf("Error building menu tree for admin ID %d: %v", Id, err)
		return nil, err
	}
	return leftMenuVo, nil
}

// 当前登录用户左侧菜单级列表(废弃)
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
