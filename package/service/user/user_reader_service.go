package user

import (
	"EduCRM/model"
	"EduCRM/package/repository"
	"EduCRM/package/store"
	"EduCRM/tools/hash"
	"EduCRM/tools/logger"
	"EduCRM/tools/response"
	"EduCRM/tools/validation"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
)

type UserReaderService struct {
	repo    *repository.Repository
	minio   *store.Store
	loggers *logger.Logger
}

func NewUserReaderService(repo *repository.Repository, minio *store.Store,
	loggers *logger.Logger) *UserReaderService {
	return &UserReaderService{repo: repo, minio: minio, loggers: loggers}
}
func (s *UserReaderService) GetUserList(role string, pagination *model.Pagination) (userList []model.User, err error) {
	if role != "all" {
		err = validation.UUIDValidation(role)
		if err != nil {
			return nil, response.ServiceError(err, codes.InvalidArgument)
		}
	}
	userList, err = s.repo.UserRepository.GetUserList(role, pagination)
	if err != nil {
		return nil, response.ServiceError(err, codes.Internal)
	}
	for i, user := range userList {
		if len(user.Photo) != 0 {
			err := s.minio.ObjectStore.ObjectExists(user.Photo)
			if err != nil {
				s.loggers.Error(err)
			} else {
				image, err := s.minio.FileLinkStore.GetImageUrl(user.Photo)
				if err != nil {
					s.loggers.Error(err)
				}
				userList[i].PhotoLink = image
			}
		}
		//userList[i].StudentGroupList, err = s.repo.GroupRepository.GetStudentByID(userList[i].ID.String())
		//if err != nil {
		//	s.loggers.Error(err)
		//}
	}
	return userList, nil
}
func (s *UserReaderService) GetUserByID(id string) (user model.User, err error) {
	user, err = s.repo.UserRepository.GetUserByID(id)
	if err != nil {
		return user, response.ServiceError(err, codes.Internal)
	}
	if len(user.Photo) != 0 {
		err = s.minio.ObjectStore.ObjectExists(user.Photo)
		if err != nil {
			s.loggers.Error(err)
		}
		image, err := s.minio.FileLinkStore.GetImageUrl(user.Photo)
		if err != nil {
			s.loggers.Error(err)
		}
		user.PhotoLink = image
	}
	return user, nil
}
func (s *UserReaderService) SignInUser(user model.SignInUser) (id, role uuid.UUID, roleTitle string, err error) {
	err = validation.ValidationStructTag(s.loggers, user)
	if err != nil {
		s.loggers.Error(err)
		return id, role, roleTitle, response.ServiceError(err, codes.InvalidArgument)
	}
	user.Password = hash.GeneratePasswordHash(user.Password)
	id, role, err = s.repo.UserRepository.SignInUser(user)
	if err != nil {
		s.loggers.Error(err)
		return id, role, roleTitle, response.ServiceError(err, codes.NotFound)
	}
	return id, role, roleTitle, nil
}
