/**
 * 路由配置
 */
package routers

import (
	config "gin-api/config/yaml_config"
	"gin-api/controllers"
	"gin-api/middlewares"
	"net/http"

	_ "gin-api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	router := gin.New()
	// 宕机恢复
	router.Use(gin.Recovery())
	// 跨域中间件
	router.Use(middlewares.Cors())
	// 图片访问路径静态文件夹可以直接访问
	router.StaticFS(config.Cfg.Upload.ImagePath, http.Dir(config.Cfg.Upload.ImagePath))
	// 日志中间件
	router.Use(middlewares.Logger())
	// 路由注册
	register(router)
	return router
}

// register 路由注册
func register(router *gin.Engine) {
	// 接口文档路径
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 验证码接口
	router.GET("/api/captcha", controllers.Captcha)
	// 登录接口
	router.POST("/api/login", controllers.Login)
	// jwt鉴权接口
	jwt := router.Group("/api", middlewares.AuthMiddleware(), middlewares.LogMiddleware())
	{
		// 用户接口
		jwt.POST("/admin/add", controllers.CreateSysAdmin)
		// jwt.GET("/admin/info", controller.GetSysAdminInfo)
		// jwt.PUT("/admin/update", controller.UpdateSysAdmin)
		// jwt.DELETE("/admin/delete", controller.DeleteSysAdminById)
		// jwt.PUT("/admin/updateStatus", controller.UpdateSysAdminStatus)
		// jwt.PUT("/admin/updatePassword", controller.ResetSysAdminPassword)
		// jwt.PUT("/admin/updatePersonal", controller.UpdatePersonal)
		// jwt.PUT("/admin/updatePersonalPassword", controller.UpdatePersonalPassword)
		// jwt.GET("/admin/list", controller.GetSysAdminList)

		// // 岗位接口
		jwt.GET("/post/list", controllers.GetSysPostList)
		jwt.POST("/post/add", controllers.AddSysPost)
		jwt.PUT("/post/update", controllers.UpdateSysPost)
		jwt.GET("/post/info", controllers.GetSysPostDetail)
		// jwt.DELETE("/post/delete", controller.DeleteSysPostById)
		// jwt.DELETE("/post/batch/delete", controller.BatchDeleteSysPost)
		// jwt.PUT("/post/updateStatus", controller.UpdateSysPostStatus)
		// jwt.GET("/post/vo/list", controller.QuerySysPostVoList)

		// // 部门接口
		// jwt.POST("/dept/add", controller.CreateSysDept)
		// jwt.GET("/dept/vo/list", controller.QuerySysDeptVoList)
		// jwt.GET("/dept/info", controller.GetSysDeptById)
		// jwt.PUT("/dept/update", controller.UpdateSysDept)
		// jwt.DELETE("/dept/delete", controller.DeleteSysDeptById)
		// jwt.GET("/dept/list", controller.GetSysDeptList)

		// // 菜单接口
		// jwt.GET("/menu/vo/list", controller.QuerySysMenuVoList)
		// jwt.POST("/menu/add", controller.CreateSysMenu)
		// jwt.GET("/menu/info", controller.GetSysMenu)
		// jwt.PUT("/menu/update", controller.UpdateSysMenu)
		// jwt.DELETE("/menu/delete", controller.DeleteSysRoleMenu)
		// jwt.GET("/menu/list", controller.GetSysMenuList)

		// // 角色接口
		// jwt.POST("/role/add", controller.CreateSysRole)
		// jwt.GET("/role/info", controller.GetSysRoleById)
		// jwt.PUT("/role/update", controller.UpdateSysRole)
		// jwt.DELETE("/role/delete", controller.DeleteSysRoleById)
		// jwt.PUT("/role/updateStatus", controller.UpdateSysRoleStatus)
		// jwt.GET("/role/list", controller.GetSysRoleList)
		// jwt.GET("/role/vo/list", controller.QuerySysRoleVoList)
		// jwt.GET("/role/vo/idList", controller.QueryRoleMenuIdList)
		// jwt.PUT("/role/assignPermissions", controller.AssignPermissions)

		// // 登录日志
		// jwt.GET("/sysLoginInfo/list", controller.GetSysLoginInfoList)
		// jwt.DELETE("/sysLoginInfo/batch/delete", controller.BatchDeleteSysLoginInfo)
		// jwt.DELETE("/sysLoginInfo/delete", controller.DeleteSysLoginInfoById)
		// jwt.DELETE("/sysLoginInfo/clean", controller.CleanSysLoginInfo)

		// // 操作日志
		// jwt.GET("/sysOperationLog/list", controller.GetSysOperationLogList)
		// jwt.DELETE("/sysOperationLog/delete", controller.DeleteSysOperationLogById)
		// jwt.DELETE("/sysOperationLog/batch/delete", controller.BatchDeleteSysOperationLog)
		// jwt.DELETE("/sysOperationLog/clean", controller.CleanSysOperationLog)

		// // 上传
		// jwt.POST("/upload", controller.Upload)
	}
}
