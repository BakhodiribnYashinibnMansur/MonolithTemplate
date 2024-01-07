package service

import (
	"EduCRM/package/repository"
	"EduCRM/package/service/minio"
	"EduCRM/package/service/user"
	"EduCRM/package/store"
	"EduCRM/tools/logger"
)

type Service struct {
	MinioService *minio.MinioService
	UserService  *user.UserService
}

func NewService(repos *repository.Repository, store *store.Store, loggers *logger.Logger) *Service {
	return &Service{
		MinioService: minio.NewMinioService(store, loggers),
		UserService:  user.NewUserService(repos, store, loggers),
	}
}
