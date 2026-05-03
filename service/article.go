package service

import (
	"blog-2026-server/model"
	"blog-2026-server/utils"
)

type ArticleService struct{}

func (s *ArticleService) GetList(current, size int) (model.PageResult[model.ArticleVO], error) {
	var count int64
	var results []model.ArticleVO

	// 1. 统计总数
	if err := utils.DB.Model(&model.Article{}).Count(&count).Error; err != nil {
		return model.PageResult[model.ArticleVO]{}, err
	}

	// 2. 只有有数据时才查询详情
	if count > 0 {
		err := utils.DB.Table("article a").
			Select("a.*, t.name as type_name").
			Joins("left join type t on a.type_id = t.id").
			Order("a.create_time desc").
			Offset((current - 1) * size).
			Limit(size).
			Scan(&results).Error

		if err != nil {
			return model.PageResult[model.ArticleVO]{}, err
		}
	}

	// 3. 组装返回结果
	return model.PageResult[model.ArticleVO]{
		List:    results,
		Total:   count,
		Current: current,
		Size:    size,
		HasMore: int64(current*size) < count,
	}, nil
}

func (s *ArticleService) GetOne(id int) (model.ArticleVO, error) {
	var result model.ArticleVO
	err := utils.DB.Table("article a").
		Select("a.*, t.name as type_name").
		Joins("left join type t on a.type_id::int = t.id").
		Where("a.id = ?", id).
		Scan(&result).Error
	return result, err
}

func (s *ArticleService) Save(article *model.Article) error {
	// 自动生成摘要：取前 120 个字
	article.Description = utils.MdGetSummary(article.Content)
	// Save 成功后，GORM 会自动把数据库生成的 ID、CreatedAt 等回填到 article 指针指向的内存，我们只需要解引用，返回这个结构体的值即可
	return utils.DB.Save(article).Error
}

func (s *ArticleService) Update(article *model.Article) error {
	return utils.DB.Model(article).Updates(article).Error
}

func (s *ArticleService) Delete(id int) error {
	return utils.DB.Delete(&model.Article{}, id).Error
}
