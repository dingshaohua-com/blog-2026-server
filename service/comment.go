package service

import (
	"blog-2026-server/model"
	"blog-2026-server/utils"
)

type CommentService struct{}

func (s *CommentService) GetCommentList(articleID int, current, size int) error {
	var comments []model.Comment
	// 1. 预加载 (Preload) 相当于自动帮你处理了 IN 查询和内存组装
	// 2. 只查一级评论 (reply_cm_id IS NULL)
	err := utils.DB.Offset((current-1)*size).Limit(size).
		Preload("Children"). // 自动查出对应的子评论
		Where("reply_article_id = ? AND reply_cm_id IS NULL", articleID).
		Order("create_time DESC").
		Find(&comments).Error
	return err
}
