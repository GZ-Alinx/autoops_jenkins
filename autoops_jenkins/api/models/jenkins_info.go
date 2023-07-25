package models

import (
	"time"
)

// Jenkins信息表字段
type JenkinsInfo struct {
	ID             uint   `json:"id" gorm:"primaryKey"`
	Url            string `json:"url" gorm:"url"`
	Login_user     string `json:"login_user" gorm:"login_user"`
	Login_password string `json:"login_password" gorm:"login_password"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// Jenkins信息表
func (JenkinsInfo) TableName() string {
	return "jenkins_info"
}
