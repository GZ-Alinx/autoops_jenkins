package router

import (
	"autoops_jenkins/api/controller"
	"autoops_jenkins/middlewares"

	"github.com/gin-gonic/gin"
)

func JenkinsRouter(Router *gin.RouterGroup) {
	// 访问 http://127.0.0.1:8080/v1/user/list
	User := Router.Group("jenkins")
	{
		User.POST("add", middlewares.JWTAuth(), controller.JenkinsAdd)
		User.POST("get", middlewares.JWTAuth(), controller.GetJenkins)
		User.POST("delete", middlewares.JWTAuth(), controller.DeleteJenkins)
		User.POST("update", middlewares.JWTAuth(), controller.UpdateJenkins)
	}
}
