package models

// 系统核心操作日志记录表
type System_log struct {
	DateTime string `json:"date_time" gorm:"date_time"`
	Action   string `json:"action" gorm:"action"`
}

func (System_log) TableName() string {
	return "system_log"
}
