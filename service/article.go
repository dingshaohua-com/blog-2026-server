package service

import (
	"blog-2026-server/model"
	"blog-2026-server/utils"
)

type ArticleService struct{}

func (s *ArticleService) GetList() ([]model.ArticleVO, error) {
	var results []model.ArticleVO
	err := utils.DB.Table("article a"). // 1. 指定主表（起别名 a）
						Select("a.*, t.name as type_name").            // 2. 手动挑选字段，注意别名要对应 VO 的 tag
						Joins("left join type t on a.type_id = t.id"). // 3. 手动写 Join 关联
						Scan(&results).Error                           // 4. 将结果“扫描”进 VO 数组
	return results, err
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
