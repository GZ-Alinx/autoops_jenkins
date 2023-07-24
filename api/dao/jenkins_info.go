package dao

import (
	"autoops_jenkins/api/forms"
	"autoops_jenkins/api/models"
	"autoops_jenkins/global"
	"fmt"

	"go.uber.org/zap"
)

func SetJenkins_Info(Jenkins_form *forms.Jenkins_form) (bool, string) {
	// 数据组装
	var Jenkins models.JenkinsInfo
	Jenkins.Login_user = Jenkins_form.Login_user
	Jenkins.Login_password = Jenkins_form.Login_password
	Jenkins.Url = Jenkins_form.Url

	// 数据处理
	if ok, response := JenkinCheckIsExistURL(Jenkins_form.Url); ok {
		return false, response
	}
	result := global.DB.Create(&Jenkins)
	if result.Error != nil {
		zap.L().Error("Jenkins添加失败")
		return false, "添加失败"
	}
	zap.L().Info("用户添加成功")
	return true, "添加成功"

}

// 检查是否存在
func JenkinCheckIsExistURL(URL string) (bool, string) {
	// 查询单个用户
	var jenkins models.JenkinsInfo
	if err := global.DB.Where("url = ?", URL).First(&jenkins).Error; err != nil {
		fmt.Println("Failed to fetch jenkins:", err)
		return false, "查询错误"
	}
	fmt.Println("查询到的jenkins信息：", jenkins)
	if len(jenkins.Url) < 1 {
		return false, "地址不存在"
	}
	return false, "地址已存在"
}

// 删除jenkins
func DeleteJenkins(url string) (bool, string) {
	var jenkins models.JenkinsInfo
	Notexist, response := JenkinCheckIsExistURL(url)
	if Notexist {
		zap.L().Info("jenkins配置不存在,无法执行删除动作")
		return false, response
	} else {
		response := global.DB.Where("url = ?", url).Delete(&jenkins)
		if response.RowsAffected > 1 {
			zap.L().Info("删除失败")
			return true, "删除失败"
		} else {
			zap.L().Info("删除成功")
			return false, "删除成功"
		}
	}
}

// jenkins获取(单个)
func Getjenkins(url string) (*models.JenkinsInfo, string) {
	var jenkins models.JenkinsInfo
	NotExist, response := JenkinCheckIsExistURL(url)
	if NotExist {
		zap.L().Info("查询的jenkins不存在")
		return &jenkins, response
	} else {
		rows := global.DB.Where("url = ?", url).Find(&jenkins)
		if rows.RowsAffected < 1 {
			zap.L().Info("查询的jenkins不存在")
			return &jenkins, response
		} else {
			return &jenkins, "查询成功"
		}
	}
}

//  jenkins job管理数据操作 不对接数据库 直接操作jenkins服务器获取
