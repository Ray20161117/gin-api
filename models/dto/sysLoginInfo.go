/**
 * 登录日志 数据层
 */
package dto

import (
	"gin-api/common/utils"
	"gin-api/models/entity"
	"gin-api/pkg/db"
	"time"
)

// 新增登录日志
func CreateSysLoginInfo(username, ipAddress, loginLocation, browser, os, message string, loginStatus int) {
	sysLoginInfo := entity.SysLoginInfo{
		Username:      username,
		IpAddress:     ipAddress,
		LoginLocation: loginLocation,
		Browser:       browser,
		Os:            os,
		Message:       message,
		LoginStatus:   loginStatus,
		LoginTime:     utils.HTime{Time: time.Now()},
	}
	db.Db.Save(&sysLoginInfo)
}
