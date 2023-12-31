package router

import (
	"autoops_jenkins/api/controller"
	"autoops_jenkins/middlewares"

	"github.com/gin-gonic/gin"
)

func DataRouter(Router *gin.RouterGroup) {
	// 访问 http://127.0.0.1:8080/v1/user/list
	User := Router.Group("data")
	{
		User.POST("add", middlewares.JWTAuth(), controller.DataAdd)
		User.POST("get", middlewares.JWTAuth(), controller.DataGet)
		User.POST("delete", middlewares.JWTAuth(), controller.DataDel)
		User.POST("update", middlewares.JWTAuth(), controller.DataUpdate)
	}
}
