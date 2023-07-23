package controller

import (
	"autoops_jenkins/api/dao"
	"autoops_jenkins/api/forms"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Summary 用户添加
// @Tags Jenkins管理
// @Produce application/json
// @Accept application/json
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Token
// @Param Token header string true "Insert your access Token"<Add access token here>
// @Param  forms.Jenkins_INFO body forms.Jenkins_INFO true "Jenkins信息"
// @Success 200
// @Router /v1/jenkins/add [post]
// Jenkins服务添加
func JenkinsAdd(ctx *gin.Context) {
	var jenkins forms.Jenkins_form
	if err := ctx.BindJSON(&jenkins); err != nil {
		zap.L().Info("数据格式错误")
		ctx.JSON(400, gin.H{
			"msg": "error",
		})
		return
	} else {
		if jenkins.Url == "" || jenkins.Login_user == "" || jenkins.Login_password == "" {
			zap.L().Info("数据格式错误")
			ctx.JSON(400, gin.H{
				"msg": "error",
			})
			return
		}
		// fmt.Println("加密密码:", string(result))
		ok, response := dao.SetJenkins_Info(&jenkins)
		if !ok {
			fmt.Println(response)
			zap.L().Error("", zap.String("status", response))
			ctx.JSON(402, gin.H{
				"msg": "error",
			})
			return
		}
		zap.L().Info(response)
		ctx.JSON(200, gin.H{
			"msg": response,
		})
	}
}

// @Summary 用户查询(单个)
// @Tags 用户管理
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Token
// @Param Token header string true "Insert your access Token"<Add access token here>
// @Param username body forms.UsernameForm true  "用户信息"
// @Success 200
// @Router /v1/jenkins/get [post]
// 用户查询(单个)
func GetJenkins(ctx *gin.Context) {
	var Jenkins forms.Jenkins_Get
	if err := ctx.BindJSON(&Jenkins); err != nil {
		zap.L().Info("数据获取错误")
		ctx.JSON(400, gin.H{
			"msg": err,
		})
		return
	} else {
		fmt.Println("请求的Url:", Jenkins.Url)
		response, str := dao.Getjenkins(Jenkins.Url)
		if str == "jenkins不存在" {
			ctx.JSON(200, gin.H{
				"msg": str,
			})
		} else {
			ctx.JSON(200, gin.H{
				"data": response,
				"msg":  str,
			})
		}

	}
}

// @Summary 用户删除
// @Tags 用户管理
// @Produce application/json
// @Accept application/json
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Token
// @Param Token header string true "Insert your access Token"<Add access token here>
// @Param  forms.UsernameForm body forms.UsernameForm true "用户信息"
// @Success 200
// @Router /v1/user/delete [post]
// 删除用户
func DeleteJenkins(ctx *gin.Context) {
	var jenkins forms.Jenkins_Get
	if err := ctx.BindJSON(&jenkins); err != nil {
		zap.L().Info("数据获取错误")
		ctx.JSON(400, gin.H{
			"msg": err,
		})
		return
	} else {
		response, str := dao.DeleteJenkins(jenkins.Url)
		if str == "jenkins不存在" {
			ctx.JSON(200, gin.H{
				"msg": str,
			})
		} else {
			ctx.JSON(200, gin.H{
				"data": response,
				"msg":  str,
			})
		}
	}

}

// @Summary 用户修改
// @Tags 用户管理
// @Produce application/json
// @Accept application/json
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Token
// @Param Token header string true "Insert your access Token"<Add access token here>
// @Param  forms.UserInfo body forms.UserInfo true "用户信息"
// @Success 200
// @Security ApiKeyAuth
// @Router /v1/user/update [post]
// 用户修改
func UpdateJenkins(ctx *gin.Context) {
	var user forms.UserInfo
	if err := ctx.BindJSON(&user); err != nil {
		zap.L().Info("用户数据错误")
		ctx.JSON(200, gin.H{
			"msg": "用户数据错误",
		})
		return
	}

	// 密码加密
	password := GeneratePassword(user.Password)
	ok, response := dao.UpdateUser(user.Username, password)
	if ok {
		zap.L().Info("用户修改成功")
		ctx.JSON(200, gin.H{
			"msg": response,
		})
		return
	}
}
