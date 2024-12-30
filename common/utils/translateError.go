/**
 * 自定义字段错误提示信息
 * 根据需要可以自行添加其它的错误提示信息
 */
package utils

import (
	"fmt"
)

// 自定义错误消息
func TranslateError(field string, tag string, param string) string {
	switch tag {
	case "required":
		return fmt.Sprintf("%s 不能为空", field)
	case "min":
		return fmt.Sprintf("%s 长度不能小于 %s", field, param)
	case "max":
		return fmt.Sprintf("%s 长度不能大于 %s", field, param)
	case "email":
		return fmt.Sprintf("%s 请提交有效的邮箱地址", field)
	default:
		return ""
	}
}
