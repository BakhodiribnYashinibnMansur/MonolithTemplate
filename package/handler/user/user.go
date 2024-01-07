package user

import (
	"EduCRM/api/jwt"
	"EduCRM/model"
	"EduCRM/package/service"
	"EduCRM/tools/handler_func"
	"EduCRM/tools/logger"
	"EduCRM/tools/response"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var (
	userID = "id"
)

type UserEndPointHandler struct {
	service *service.Service
	loggers *logger.Logger
}

func NewUserEndPointHandler(service *service.Service,
	loggers *logger.Logger) *UserEndPointHandler {
	return &UserEndPointHandler{service: service, loggers: loggers}
}

// CreateUser
// @Description Create User
// @Summary Create User
// @Tags User
// @Accept json
// @Produce json
// @Param create body model.CreateUser true "Create User"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/user/create [post]
// @Security ApiKeyAuth
func (h *UserEndPointHandler) CreateUser(ctx *gin.Context) {
	var (
		body model.CreateUser
	)
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil, nil)
		return
	}
	_, _, err = h.service.UserService.CreateUser(body)
	if err != nil {
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.Created, nil, "created", nil)
}

// UpdateUser
// @Description Update User
// @Summary Update User
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param update body model.UpdateUser true "Update User"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/user/update/{id} [put]
// @Security ApiKeyAuth
func (h *UserEndPointHandler) UpdateUser(ctx *gin.Context) {
	var (
		body model.UpdateUser
	)
	userID, err := handler_func.GetUUIDParam(ctx, userID)
	if err != nil {
		response.HandleResponse(ctx, response.NotFound, err, nil, nil)
		return
	}
	err = ctx.ShouldBindJSON(&body)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil, nil)
		return
	}
	body.ID = userID
	_, err = h.service.UserService.UpdateUser(body)
	if err != nil {
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.OK, nil, "updated", nil)
}

// DeleteUser
// @Description Delete User
// @Summary Delete User
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/user/delete/{id} [delete]
// @Security ApiKeyAuth
func (h *UserEndPointHandler) DeleteUser(ctx *gin.Context) {
	id, err := handler_func.GetUUIDParam(ctx, userID)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil, nil)
		return
	}
	err = h.service.UserService.DeleteUser(id)
	if err != nil {
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.OK, nil, errors.New("deleted"), nil)
}

// GetUserList
// @Description Get User List
// @Summary Get User List
// @Tags User
// @Accept json
// @Produce json
// @Param pageSize query int64 true "pageSize" default(10)
// @Param page  query int64 true "page" default(1)
// @Param role  query string true "Role : all.  get all user"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/user/list [get]
// @Security ApiKeyAuth
func (h *UserEndPointHandler) GetUserList(ctx *gin.Context) {
	pagination, err := handler_func.ListPagination(ctx)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil, nil)
		return
	}
	role, err := handler_func.GetStringQuery(ctx, "role")
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil, nil)
		return
	}
	userList, err := h.service.UserService.GetUserList(role, &pagination)
	if err != nil {
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.OK, nil,
		userList,
		pagination)
}

// GetUserByID
// @Description Get User
// @Summary Get User
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/user/{id} [get]
// @Security ApiKeyAuth
func (h *UserEndPointHandler) GetUserByID(ctx *gin.Context) {
	id, err := handler_func.GetUUIDParam(ctx, userID)
	if err != nil {
		response.HandleResponse(ctx, response.NotFound, err, nil, nil)
		return
	}
	user, err := h.service.UserService.GetUserByID(id.String())
	if err != nil {
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.OK, nil, user, nil)
}

// SignInUser
// @Description Admin Sign In  User.
// @Summary Admin Sign In User
// @Tags User
// @Accept json
// @Produce json
// @Param signup body model.SignInUser true "Sign In"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/user/sign-in [post]
func (h *UserEndPointHandler) SignInUser(ctx *gin.Context) {
	var (
		body model.SignInUser
	)
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil, nil)
	}
	body.PhoneNumber = strings.TrimSpace(body.PhoneNumber)
	body.Password = strings.TrimSpace(body.Password)
	id, role, roleTitle, err := h.service.UserService.SignInUser(body)
	if err != nil {
		response.ServiceErrorConvert(ctx, err)
		return
	}
	if id == uuid.Nil || role == uuid.Nil {
		response.HandleResponse(ctx, response.BadRequest,
			errors.New("username or password is incorrect"), nil, nil)
		return
	}
	tokens, err := jwt.GenerateNewTokens(id.String(), role.String(), roleTitle)
	if err != nil {
		response.HandleResponse(ctx, response.InternalServerError, err, nil, nil)
		return
	}
	response.HandleResponse(ctx, response.OK, nil, tokens, nil)
}

// UpdateUserPassword
// @Description Update User Password
// @Summary Update User Password
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param password body model.UserPassword true "password"
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/user/update-password/{id} [put]
// @Security ApiKeyAuth
func (h *UserEndPointHandler) UpdateUserPassword(ctx *gin.Context) {
	id, err := handler_func.GetUUIDParam(ctx, userID)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil, nil)
		return
	}
	var user model.UserPassword
	err = ctx.ShouldBindJSON(&user)
	if err != nil {
		response.HandleResponse(ctx, response.BadRequest, err, nil, nil)
		return
	}
	err = h.service.UserService.UpdateUserPassword(id.String(), user.Password)
	if err != nil {
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.OK, nil, "updated", nil)
}

// GetUserMe
// @Description User Me
// @Summary User Me
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} response.ResponseModel
// @Failure 400 {object} response.ResponseModel
// @Failure 404 {object} response.ResponseModel
// @Failure 500 {object} response.ResponseModel
// @Router /api/v1/user/me [get]
// @Security ApiKeyAuth
func (h *UserEndPointHandler) GetUserMe(ctx *gin.Context) {
	id, err := handler_func.GetUserId(ctx)
	if err != nil {
		response.HandleResponse(ctx, response.NotFound, err, nil, nil)
		return
	}
	user, err := h.service.UserService.GetUserByID(id)
	if err != nil {
		response.ServiceErrorConvert(ctx, err)
		return
	}
	response.HandleResponse(ctx, response.OK, nil, user, nil)
}
