package model

// PageResult 分页统一返回结构
type PageResult[T any] struct {
	List    []T   `json:"list"`    // 数据列表
	Total   int64 `json:"total"`   // 总条数
	Current int   `json:"current"` // 当前页码
	Size    int   `json:"size"`    // 每页条数
	HasMore bool  `json:"hasMore"` // 是否有更多
}
