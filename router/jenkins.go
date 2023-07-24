package router

import (
	"autoops_jenkins/api/controller"
	"autoops_jenkins/middlewares"

	"github.com/gin-gonic/gin"
)

func JenkinsRouter(Router *gin.RouterGroup) {
	// 访问 http://127.0.0.1:8080/v1/user/list
	Jenkins := Router.Group("jenkins")
	{
		Jenkins.POST("add", middlewares.JWTAuth(), controller.JenkinsAdd)
		Jenkins.POST("get", middlewares.JWTAuth(), controller.GetJenkins)
		Jenkins.POST("delete", middlewares.JWTAuth(), controller.DeleteJenkins)
		Jenkins.POST("update", middlewares.JWTAuth(), controller.UpdateJenkins)
		Jenkins.POST("gettoken", middlewares.JWTAuth(), controller.JenkinsToekn)
		Jenkins.POST("jobadd", middlewares.JWTAuth(), controller.JenkinsJobAdd)
	}
}
