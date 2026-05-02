package model

import (
	"time"
)

// Comment 评论实体
type Comment struct {
	// ID 使用 uint 更符合 Go 的习惯
	ID             uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Content        string    `gorm:"type:text;not null" json:"content"`
	CreateTime     time.Time `gorm:"autoCreateTime" json:"create_time"`
	ReplyArticleId uint      `gorm:"index" json:"reply_article_id"`

	// 用户相关信息
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
	BlogUrl  string `json:"blog_url"`
	NickName string `json:"nick_name"`

	// 父级评论ID，使用指针类型 (*uint) 方便处理 NULL 值（即根评论）
	ReplyCmId *uint `gorm:"index" json:"reply_cm_id"`

	// --- 重点：摆到明面上的“树”结构 ---
	// 在 Java 中你是用 Map.put("children", list) 动态塞进去的
	// 在 Go 中我们直接定义好，GORM 预加载时会自动填充
	Children []Comment `gorm:"foreignKey:ReplyCmId" json:"children"`
}
