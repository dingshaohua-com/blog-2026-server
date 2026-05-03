package model

import "time"

type Article struct {
	// 必须大写开头！后面反引号里的才是数据库真实的列名
	Id          int       `gorm:"column:id" json:"id"`
	Title       string    `gorm:"column:title" json:"title"`
	Description string    `gorm:"column:description" json:"description"`
	TypeId      string    `gorm:"column:type_id" json:"typeId"`
	CreateTime  time.Time `gorm:"column:create_time" json:"createTime"`
	Content     string    `gorm:"column:content" json:"content"`
}

// ArticleVO 展示模型 (VO)，直接在下面定义，不用谢重复字段，用匿名嵌套
type ArticleVO struct {
	Article         // 组合 Article 的所有字段
	TypeName string `json:"typeName" gorm:"column:type_name"` // 扩展字段
}
