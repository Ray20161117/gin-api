/**
 * 时间工具类
 */
package utils

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type HTime struct {
	time.Time
}

var formatTime = "2006-01-02 15:04:05"

// 将时间对象格式化JSON兼容的字符串格式
func (t HTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format(formatTime))
	return []byte(formatted), nil
}

// 将JSON格式的时间字符串解析未HTime结构体
func (t *HTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+formatTime+`"`, string(data), time.Local)
	*t = HTime{Time: now}
	return
}

// 实现HTime类型与数据库的交互,实现了将 HTime 类型转换为数据库驱动可以接受的值的方法
func (t HTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// 主要功能是实现 HTime 类型与数据库的交互，具体来说，它是一个用于从数据库中扫描时间值并将其赋给 HTime 类型变量的方法。
func (t *HTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = HTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
