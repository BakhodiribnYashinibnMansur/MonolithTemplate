package user

import (
	"EduCRM/package/service"
	"EduCRM/tools/logger"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserEndPoint
}
type UserEndPoint interface {
	SignInUser(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
	GetUserList(ctx *gin.Context)
	GetUserByID(ctx *gin.Context)
	UpdateUserPassword(ctx *gin.Context)
	GetUserMe(ctx *gin.Context)
}

func NewUserHandler(service *service.Service,
	loggers *logger.Logger) *UserHandler {
	return &UserHandler{
		UserEndPoint: NewUserEndPointHandler(service, loggers),
	}
}
