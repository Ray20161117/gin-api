/**
 * 操作日志 数据层
 */
package dto

import (
	"gin-api/models/entity"
	"gin-api/pkg/db"
)

// 新增操作日志
func CreateSysOperationLog(log entity.SysOperationLog) {
	db.Db.Create(&log)
}
