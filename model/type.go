package model

// Type 文章分类，对应 Java 中的 Type 实体
type Type struct {
	Id          int    `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	Name        string `gorm:"column:name" json:"name"`
	Description string `gorm:"column:description" json:"description"`
}

// TableName 显式指定表名，避免 GORM 默认按复数 types 去查
func (Type) TableName() string {
	return "type"
}
