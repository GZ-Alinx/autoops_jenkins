package initialize

import (
	"autoops_jenkins/global"
	"fmt"
	"time"

	"github.com/fatih/color"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysqlDB() error {
	// 参考资料： https://github.com/go-sql-driver/mysql#dsn-data-source-name
	mysqlInfo := global.Config.MysqlInfo
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlInfo.User,
		mysqlInfo.PassWord,
		mysqlInfo.Host,
		mysqlInfo.Port,
		mysqlInfo.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		zap.L().Error("数据库链接错误，err:", zap.String("err", err.Error()))
		return err
	}
	sqlDB, err := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(50)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(200)
	// SetConnMaxLifetime 设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)
	// db.AutoMigrate(models.User{}) 自动创建数据表
	global.DB = db
	color.Green("MYSQL数据库初始化成功")
	return nil
}
