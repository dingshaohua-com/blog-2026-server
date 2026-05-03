package service

import (
	"blog-2026-server/model"
	"blog-2026-server/utils"
)

type TypeService struct{}

// GetList 分页获取分类列表
func (s *TypeService) GetList(current, size int) (model.PageResult[model.Type], error) {
	var count int64
	var list []model.Type

	db := utils.DB.Model(&model.Type{})

	if err := db.Count(&count).Error; err != nil {
		return model.PageResult[model.Type]{}, err
	}

	if count > 0 {
		if err := db.Order("id desc").
			Offset((current - 1) * size).
			Limit(size).
			Find(&list).Error; err != nil {
			return model.PageResult[model.Type]{}, err
		}
	}

	return model.PageResult[model.Type]{
		List:    list,
		Total:   count,
		Current: current,
		Size:    size,
		HasMore: int64(current*size) < count,
	}, nil
}

// GetAll 获取所有分类（不分页，常用于下拉框）
func (s *TypeService) GetAll() ([]model.Type, error) {
	var list []model.Type
	err := utils.DB.Order("id asc").Find(&list).Error
	return list, err
}

// GetOne 根据 ID 查询单个分类
func (s *TypeService) GetOne(id int) (model.Type, error) {
	var t model.Type
	err := utils.DB.First(&t, id).Error
	return t, err
}

// Save 新增分类
func (s *TypeService) Save(t *model.Type) error {
	return utils.DB.Create(t).Error
}

// Update 根据 ID 更新分类信息
func (s *TypeService) Update(t *model.Type) error {
	return utils.DB.Model(t).Updates(t).Error
}

// Delete 删除分类
func (s *TypeService) Delete(id int) error {
	return utils.DB.Delete(&model.Type{}, id).Error
}
