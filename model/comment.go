package model

import (
	"time"
)

// Comment 评论实体，对应 Java 中的 Comment
type Comment struct {
	Id             int       `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	Content        string    `gorm:"column:content" json:"content"`
	CreateTime     time.Time `gorm:"column:create_time;autoCreateTime" json:"createTime"`
	ReplyArticleId int       `gorm:"column:reply_article_id;index" json:"replyArticleId"`

	Avatar   string `gorm:"column:avatar" json:"avatar"`
	Email    string `gorm:"column:email" json:"email"`
	BlogUrl  string `gorm:"column:blog_url" json:"blogUrl"`
	NickName string `gorm:"column:nick_name" json:"nickName"`

	// 父级评论 ID，使用指针以表达 NULL（即根评论）
	ReplyCmId *int `gorm:"column:reply_cm_id;index" json:"replyCmId"`

	// 子评论树：Java 端用 Map 动态塞进 children，这里直接显式声明
	// GORM 在使用 Preload("Children") 时会自动按 reply_cm_id 关联回填
	Children []Comment `gorm:"foreignKey:ReplyCmId;references:Id" json:"children"`
}
