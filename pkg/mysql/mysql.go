/**
* Mysql 数据库连接
 */

package mysql

import (
	"fmt"
	config "gin-api/config/yaml_config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

func InitMysql() error {
	var err error
	var DbConfig = config.Cfg.Database
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		DbConfig.Username,
		DbConfig.Password,
		DbConfig.Host,
		DbConfig.Port,
		DbConfig.Database,
		DbConfig.Charset)
	Db, err = gorm.Open(mysql.Open(url), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}
	if Db.Error != nil {
		panic(Db.Error)
	}
	sqlDB, _ := Db.DB()
	sqlDB.SetMaxIdleConns(DbConfig.MaxIdle)
	sqlDB.SetMaxOpenConns(DbConfig.MaxOpen)
	return nil
}
