package model

import (
	"gorm.io/datatypes"
)

// --- 列表项子结构定义 ---

// FriendItem 友情链接
type FriendItem struct {
	Name   string `json:"name" example:"丁少华"`
	Dec    string `json:"dec" example:"站主的大神男友"`
	Url    string `json:"url" example:"https://www.cnblogs.com/dingshaohua/"`
	Avatar string `json:"avatar" example:"https://..."`
}

// PluginItem 插件信息
type PluginItem struct {
	Name  string `json:"name" example:"md-editor-rt"`
	Dec   string `json:"dec" example:"基于React的markdown..."`
	Tab   string `json:"tab" example:"React"`
	Url   string `json:"url" example:"https://..."`
	Color string `json:"color" example:"orange"`
}

// WebsiteItem 常用网站
type WebsiteItem struct {
	Name  string `json:"name" example:"PDF转Word转换器"`
	Dec   string `json:"dec" example:"完美而迅速地将PDF转换..."`
	Tab   string `json:"tab" example:"工具"`
	Url   string `json:"url" example:"https://..."`
	Color string `json:"color" example:"#87d068"`
}

// --- JSON 核心内容结构 ---

// AppContent 对应 JSONB 内部的具体结构
type AppContent struct {
	// 用户基础信息
	Nickname string `json:"nickname" example:"花贝"`
	Dec      string `json:"dec" example:"三分钟热度的实践者"`
	Pwd      string `json:"pwd" example:"123456"`
	Role     string `json:"role" example:"root"`
	Email    string `json:"email" example:"869043928@qq.com"`
	Avatar   string `json:"avatar" example:"https://..."`

	// 站点元数据
	StationName        string `json:"stationName" example:"花贝の碎碎念"`
	StationDec         string `json:"stationDec" example:"The tough road..."`
	StationPutOnRecord string `json:"stationPutOnRecord" example:"京ICP备2021005859号"`

	// 动态列表数据
	Friend  []FriendItem  `json:"friend"`
	Plugins []PluginItem  `json:"plugins"`
	Website []WebsiteItem `json:"website"`
}

// --- 数据库主模型 ---

type AppInfo struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Version string `json:"version" example:"0.1"`
	// 使用泛型 JSONType 绑定，实现自动序列化/反序列化
	Content datatypes.JSONType[AppContent] `json:"content"`
}
