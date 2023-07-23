package main

import (
	"autoops_jenkins/api/controller"
	"autoops_jenkins/api/models"
	"autoops_jenkins/global"

	"autoops_jenkins/initialize"
	"fmt"

	"github.com/fatih/color"
	"go.uber.org/zap"
)

func main() {
	// 配置初始化
	initialize.InitConfig()
	// MYSQL数据库初始化
	if err := initialize.InitMysqlDB(); err != nil {
		color.Red("数据库初始化异常")
		panic(err)
	}

	// REDIS数据库初始化
	if err := initialize.InitRedis(); err != nil {
		color.Red("数据库初始化异常")
		panic(err)
	}

	// 日志配置初始化
	initialize.InitLogger()

	// 路由配置初始化
	Router := initialize.Routers()

	// 表结构迁移
	// 迁移数据表（仅在第一次运行时执行）
	global.DB.AutoMigrate(&models.User{})
	global.DB.AutoMigrate(&models.Data{})
	global.DB.AutoMigrate(&models.JenkinsInfo{})
	global.DB.AutoMigrate(&models.Host{})
	global.DB.AutoMigrate(&models.System_log{})

	// 管理员用户添加
	ret := controller.UserAddAdmin()
	fmt.Println(ret)

	// 服务启动
	err := Router.Run(fmt.Sprintf(":%d", global.Config.Port))
	if err != nil {
		zap.L().Info("this is hello func", zap.String("error", "启动错误"))
	}
	color.Green("启动成功")
}
