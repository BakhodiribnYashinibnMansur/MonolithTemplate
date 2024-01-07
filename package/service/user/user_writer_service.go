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

var (
// groupID    = "groupID"
// newGroupID = "newGroupID"
// oldGroupID   = "oldGroupID"
// teacherID    = "teacherID"
// newTeacherID = "newTeacherID"
// oldTeacherID = "oldTeacherID"
)

type UserWriterService struct {
	repo    *repository.Repository
	minio   *store.Store
	loggers *logger.Logger
}

func NewUserWriterService(repo *repository.Repository, minio *store.Store,
	loggers *logger.Logger) *UserWriterService {
	return &UserWriterService{repo: repo, minio: minio, loggers: loggers}
}
func (s *UserWriterService) CreateUser(user model.CreateUser) (id uuid.UUID, role string, err error) {
	err = validation.ValidationStructTag(s.loggers, user)
	if err != nil {
		return id, role, response.ServiceError(err, codes.InvalidArgument)
	}
	user.Password = hash.GeneratePasswordHash(user.Password)
	userID, err := s.repo.UserRepository.CreateUser(user)
	if err != nil {
		return id, role, response.ServiceError(err, codes.Internal)
	}
	return userID, user.RoleID, nil
}
func (s *UserWriterService) UpdateUser(user model.UpdateUser) (role string, err error) {
	err = validation.ValidationStructTag(s.loggers, user)
	if err != nil {
		return role, response.ServiceError(err, codes.InvalidArgument)
	}
	err = s.repo.UserRepository.UpdateUser(user)
	if err != nil {
		return role, response.ServiceError(err, codes.Internal)
	}
	return user.RoleID, nil
}
func (s *UserWriterService) DeleteUser(id uuid.UUID) (err error) {
	err = s.repo.UserRepository.DeleteUser(id.String())
	if err != nil {
		return response.ServiceError(err, codes.Internal)
	}
	// err = s.repo.GroupRepository.DeleteUserAllGroupEnrollment(id)
	// if err != nil {
	// 	return ServiceErrorHandler(err, codes.Internal)
	// }
	//delete from auth_account
	return nil
}
func (s *UserWriterService) UpdateUserPassword(id, password string) error {
	//err := validation.ValidatePassword(password, "password")
	//if err != nil {
	//	return response.ServiceError(err, codes.InvalidArgument)
	//}
	password = hash.GeneratePasswordHash(password)
	err := s.repo.UserRepository.UpdateUserPassword(id, password)
	if err != nil {
		return response.ServiceError(err, codes.Internal)
	}
	return nil
}
