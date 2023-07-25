package dao

import (
	"autoops_jenkins/api/models"
	"autoops_jenkins/global"
	"fmt"

	"go.uber.org/zap"
)

// 通过用户名查询用户信息
// func GetUserByUsername(username string) models.User {
// 	user := models.User{}
// 	global.DB.Find(&user, global.DB.Where("username = ?", username))
// 	return user
// }

// 检查用户是否存在
func GetUserByUsername(username string) models.User {
	// 查询单个用户
	var user models.User
	if err := global.DB.Where("username = ?", username).First(&user).Error; err != nil {
		fmt.Println("Failed to fetch user:", err)
		return user
	}

	if user.Username == "" {
		global.Lg.Error("", zap.String("status", "用户不存在"))
		return user
	}
	return user
}
