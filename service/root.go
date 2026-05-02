package service

import (
	"blog-2026-server/model"
	"blog-2026-server/utils"
)

type RootService struct{}

func (s *RootService) GetAppInfo() (*model.AppInfo, error) {
	var apps model.AppInfo
	err := utils.DB.First(&apps, 1).Error
	return &apps, err
}
