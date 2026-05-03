package model

import (
	"time"
)

// Mood 对应 Java 的 Mood 实体
// 使用 swaggo 标签配合 Swagger 文档
type Mood struct {
	ID         int       `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	Content    string    `gorm:"column:content" json:"content"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime" json:"createTime"`
}
