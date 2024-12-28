/**
 * @Author: Jiayo
 * @Date: 2021-08-18 14:50:23
 * @LastEditTime: 2021-08-18 14:50:23
 * @LastEditors: Please set LastEditors
 */
package main

import (
	"context"
	config "gin-api/config/yaml_config"
	"gin-api/pkg/db"
	"gin-api/pkg/log"
	"gin-api/pkg/redis"
	"gin-api/routers"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

// @title 通用后台管理系统
// @version 1.0
// @description 后台管理系统API接口文档
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// 加载日志log
	log := log.Log()
	// 设置启动模式
	gin.SetMode(config.Cfg.Server.Model)
	// 初始化路由
	router := routers.InitRouter()
	srv := &http.Server{
		Addr:    config.Cfg.Server.Address,
		Handler: router,
	}
	// 启动服务
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Info("listen:", err)
		}
		log.Info("listen:", config.Cfg.Server.Address)
	}()
	quit := make(chan os.Signal, 1)
	// 监听消息
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Info("Shutdown Server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Info("Server Shutdown:", err)
	}
	log.Info("Server exiting")
}

// 初始化连接
func init() {
	// mysql
	db.InitMysql()
	// redis
	redis.InitRedis()
}
