package user

import (
	"EduCRM/model"
	"EduCRM/package/repository"
	"EduCRM/package/store"
	"EduCRM/tools/logger"

	"github.com/google/uuid"
)

type UserService struct {
	User
}
type User struct {
	UserReader
	UserWriter
}
type UserWriter interface {
	CreateUser(user model.CreateUser) (id uuid.UUID, role string, err error)
	UpdateUser(user model.UpdateUser) (role string, err error)
	DeleteUser(id uuid.UUID) (err error)
	UpdateUserPassword(id, password string) error
}
type UserReader interface {
	GetUserList(role string, pagination *model.Pagination) (userList []model.User, err error)
	GetUserByID(id string) (user model.User, err error)
	SignInUser(user model.SignInUser) (id, role uuid.UUID, roleTitle string, err error)
}

func NewUserService(repos *repository.Repository, store *store.Store,
	loggers *logger.Logger) *UserService {
	return &UserService{
		User: User{
			UserReader: NewUserReaderService(repos, store, loggers),
			UserWriter: NewUserWriterService(repos, store, loggers),
		},
	}
}
