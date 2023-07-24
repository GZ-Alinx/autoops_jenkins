package controller

import (
	"autoops_jenkins/api/dao"
	"autoops_jenkins/global"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	jenkinsURL       = "http://your-jenkins-server-url"
	jenkinsUsername  = "your-jenkins-username"
	jenkinsAPIToken  = "your-jenkins-api-token"
	jenkinsJobFolder = "your-job-folder"                                // 可选的，如果需要将流水线放在特定的文件夹中
	gitLabRepoURL    = "https://gitlab.com/your-username/your-repo.git" // 外部GitLab仓库URL
	gitLabBranch     = "main"                                           // 外部GitLab仓库的分支名称
)

type JenkinsUser struct {
	ID       string `json:"id"`
	Fullname string `json:"fullName"`
	Token    string `json:"token"`
}

func JenkinsToekn(ctx *gin.Context) {
	url, user, passwd := JenkinsTokenSet()
	fmt.Println(url, user, passwd)
	ctx.JSON(200, gin.H{
		"msg":    "success",
		"user":   user,
		"passwd": passwd,
		"url":    url,
	})
}

// 获取jenkinsToken
func JenkinsTokenSet() (url, username, password string) {
	// Jenkins 实例的 URL 和用户名密码
	data, reponse := dao.Getjenkins("http://121.37.204.75")
	fmt.Println("status", reponse)
	global.Lg.Info("", zap.String("status", reponse))
	return data.Url, data.Login_user, data.Login_password
}

func Create_Job() {
	jenkinsURL, username, password := JenkinsTokenSet()
	// 构建Jenkins Job配置
	jobConfig := map[string]interface{}{
		"branchSources": []map[string]interface{}{
			{
				"source": map[string]interface{}{
					"remote":        gitLabRepoURL,
					"credentialsId": "jenkins-gitlab-credentials", // Jenkins中GitLab凭证的ID
				},
				"strategy": map[string]interface{}{
					"type":  "honorRefs",
					"regex": fmt.Sprintf("(%s)", gitLabBranch),
				},
			},
		},
		"jenkinsfilePath": "Jenkinsfile", // 指定外部GitLab仓库中的Jenkinsfile路径
	}

	// 将Jenkins Job配置转换为JSON格式
	jobConfigJSON, err := json.Marshal(jobConfig)
	if err != nil {
		fmt.Println("JSON编码失败:", err)
		return
	}

	// 发送HTTP请求创建Jenkins流水线
	url := jenkinsURL + "/createItem"
	if jenkinsJobFolder != "" {
		url = jenkinsURL + "/job/" + jenkinsJobFolder + "/createItem"
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jobConfigJSON))
	if err != nil {
		fmt.Println("创建HTTP请求失败:", err)
		return
	}

	req.SetBasicAuth(username, password)
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送HTTP请求失败:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("流水线已成功创建.")
}

// jenkins job manager
func JenkinsJobAdd(ctx *gin.Context) {
	Create_Job()
	ctx.JSON(200, gin.H{
		"status": "success",
	})
}

// // jenkins 删除
// func JenkinsJobDel(ctx *gin.Context) {

// }

// // jenkins 获取
// func JenkinsGet(ctx *gin.Context) {

// }

// // jenkins更新
// func JenkinsJobUpdate(ctx *gin.Context) {

// }
