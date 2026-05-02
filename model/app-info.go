package model

import (
	"gorm.io/datatypes"
)

// AppContent 定义 JSONB 内部的具体结构
type AppContent struct {
	Pwd      string `json:"pwd"`
	Role     string `json:"role"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	Nickname string `json:"nickname"`
}

type AppInfo struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Version string `json:"version"`
	// 使用 JSONType 绑定上面的结构体
	Content datatypes.JSONType[AppContent] `json:"content"`
}
