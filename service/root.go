package service

import (
	"blog-server/model"
	"blog-server/utils"
)

type RootService struct{}

func (s *RootService) GetAppInfo() error {
	var apps model.AppInfo
	return utils.DB.First(&apps, 1).Error
}
