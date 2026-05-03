package service

import (
	"blog-2026-server/model"
	"blog-2026-server/utils"
)

type MoodService struct{}

// Create 添加心情
func (s *MoodService) Create(mood model.Mood) error {
	return utils.DB.Create(&mood).Error
}

// Delete 删除心情
func (s *MoodService) Delete(id int) error {
	return utils.DB.Delete(&model.Mood{}, id).Error
}

// Update 更新心情（根据 ID 更新内容）
func (s *MoodService) Update(mood model.Mood) error {
	// Select("Content") 表示只更新内容字段，保护 create_time 不被篡改
	return utils.DB.Model(&mood).Select("Content").Updates(mood).Error
}

// GetList 分页获取
func (s *MoodService) GetList(current, size int) (model.PageResult[model.Mood], error) {
	var count int64
	var list []model.Mood

	db := utils.DB.Model(&model.Mood{})

	// 1. 统计总数
	if err := db.Count(&count).Error; err != nil {
		return model.PageResult[model.Mood]{}, err
	}

	// 2. 查询数据
	err := db.Order("create_time desc").
		Offset((current - 1) * size).
		Limit(size).
		Find(&list).Error

	return model.PageResult[model.Mood]{
		List:    list,
		Total:   count,
		Current: current,
		Size:    size,
		HasMore: int64(current*size) < count,
	}, err
}
