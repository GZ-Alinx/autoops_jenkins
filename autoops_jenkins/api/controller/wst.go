package controller

import (
	"autoops_jenkins/api/forms"
	"fmt"
	"github.com/gin-gonic/gin"
)

func WEBSET(ctx *gin.Context) {
	info := forms.WSHost{
		Username:  "root",
		Password:  "",
		Port:      22,
		Ipaddress: "42.193.154.221",
	}
	fmt.Println(info)
	ctx.HTML(200, "/static/front/inedex.html", gin.H{})
}
