package service

import (
	"blog-2026-server/model"
	"blog-2026-server/utils"

	"gorm.io/gorm"
)

type CommentService struct{}

// GetListByArticle 分页获取某篇文章下的根评论（reply_cm_id IS NULL）
// 同时通过 Preload 自动把子评论挂到 Children 字段
func (s *CommentService) GetListByArticle(articleId, current, size int) (model.PageResult[model.Comment], error) {
	var count int64
	var list []model.Comment

	// 仅对“根评论”做分页统计，子评论跟随父级一起加载
	db := utils.DB.Model(&model.Comment{}).
		Where("reply_article_id = ? AND reply_cm_id IS NULL", articleId)

	if err := db.Count(&count).Error; err != nil {
		return model.PageResult[model.Comment]{}, err
	}

	if count > 0 {
		err := db.
			Preload("Children"). // 自动查出子评论
			Order("create_time desc").
			Offset((current - 1) * size).
			Limit(size).
			Find(&list).Error
		if err != nil {
			return model.PageResult[model.Comment]{}, err
		}
	}

	return model.PageResult[model.Comment]{
		List:    list,
		Total:   count,
		Current: current,
		Size:    size,
		HasMore: int64(current*size) < count,
	}, nil
}

// GetOne 根据 ID 获取评论详情（含子评论）
func (s *CommentService) GetOne(id int) (model.Comment, error) {
	var c model.Comment
	err := utils.DB.Preload("Children").First(&c, id).Error
	return c, err
}

// Save 新增评论（包括根评论与回复评论）
func (s *CommentService) Save(c *model.Comment) error {
	return utils.DB.Create(c).Error
}

// Update 根据 ID 更新评论内容等信息
func (s *CommentService) Update(c *model.Comment) error {
	return utils.DB.Model(c).Updates(c).Error
}

// Delete 删除评论；为防止出现“悬空子评论”，同时把它的子评论一起清理
// 整个过程用事务包起来，要么全部成功要么全部回滚
func (s *CommentService) Delete(id int) error {
	return utils.DB.Transaction(func(tx *gorm.DB) error {
		// 1. 先删除挂在该评论下的子评论
		if err := tx.Where("reply_cm_id = ?", id).Delete(&model.Comment{}).Error; err != nil {
			return err
		}
		// 2. 再删除评论本身
		return tx.Delete(&model.Comment{}, id).Error
	})
}
