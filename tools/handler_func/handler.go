package handler_func

import (
	"EduCRM/config"
	"EduCRM/model"
	"EduCRM/tools/logger"
	"EduCRM/tools/response"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var (
	cfg     = config.Config()
	loggers = logger.GetLogger()
)

const (
	userIdCtx    = "userId"
	userRoleCtx  = "userRole"
	paramInvalid = "%s param is invalid"
	queryInvalid = "%s query is invalid"
	ParseDate    = "2006-01-02T15:04:05Z07:00"
	FormatDate   = "2006-01-02 15:04:05"
	page         = "page"
	pageSize     = "page-size"
)

func GetStringParam(ctx *gin.Context, query string) (param string, err error) {
	param = ctx.Param(query)
	if param == "" {
		err := fmt.Sprintf(" %s param is empty", query)
		return "", errors.New(err)
	}
	return param, nil
}
func GetNullStringParam(ctx *gin.Context, query string) (param string, err error) {
	param = ctx.Param(query)
	return param, nil
}
func GetNullStringQuery(ctx *gin.Context, query string) (param string) {
	param = ctx.Query(query)
	return param
}
func GetStringQuery(ctx *gin.Context, query string) (param string, err error) {
	param = ctx.Query(query)
	if param == "" {
		err := fmt.Sprintf(" %s param is empty", query)
		return "", errors.New(err)
	}
	return param, nil
}
func GetInt64Query(ctx *gin.Context,
	query string) (int64,
	error) {
	param := ctx.Query(query)
	if param == "" {
		return 0, nil
	}
	paramInt, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return 0, err
	}
	return paramInt, nil
}
func GetFloat64Query(ctx *gin.Context, query string) (float64,
	error) {
	param := ctx.Query(query)
	if param == "" {
		return 0, nil
	}
	paramInt, err := strconv.ParseFloat(param, 64)
	if err != nil {
		return 0, err
	}
	return paramInt, nil
}
func GetArrayStringQuery(ctx *gin.Context,
	query string) ([]string, error) {
	param := ctx.Query(query)
	if param == "" {
		return []string{}, errors.New("param is empty")
	}
	chunks := strings.Split(param, ",")
	return chunks, nil
}
func JsonUnmarshal(pointData *interface{},
	data []byte) error {
	err := json.Unmarshal(data, pointData)
	if err != nil {
		return err
	}
	return nil
}
func GetBooleanQuery(ctx *gin.Context, query string) (bool, error) {
	param := ctx.Query(query)
	if param == "" {
		err := fmt.Sprintf(paramInvalid, query)
		return false, errors.New(err)
	}
	boolVal, err := strconv.ParseBool(param)
	if err != nil {
		err := fmt.Sprintf(paramInvalid, query)
		return false, errors.New(err)
	}
	return boolVal, nil
}
func GetUUIDQuery(ctx *gin.Context, query string) (uuid.UUID, error) {
	param := ctx.Query(query)
	if param == "" {
		err := fmt.Sprintf(paramInvalid, query)
		return uuid.Nil, errors.New(err)
	}
	paramUUID, err := uuid.Parse(param)
	if err != nil {
		loggers.Error(err)
		err := fmt.Sprintf(paramInvalid, query)
		return uuid.Nil, errors.New(err)
	}
	return paramUUID, nil
}
func GetUserId(ctx *gin.Context) (string, error) {
	id, ok := ctx.Get(userIdCtx)
	if !ok {
		loggers.Error(response.ErrorUserIDInvalid.Error())
		return "", errors.New(response.ErrorUserIDInvalid.Error())
	}
	userID, ok := id.(string)
	if !ok {
		loggers.Error(response.ErrorUserIDInvalid.Error())
		return "", errors.New(response.ErrorUserIDInvalid.Error())
	}
	_, err := uuid.Parse(userID)
	if err != nil {
		loggers.Error(err)
		return "", err
	}
	return userID, nil
}
func GetUserUUID(ctx *gin.Context) (uuid.UUID, error) {
	id, ok := ctx.Get(userIdCtx)
	if !ok {
		loggers.Error(response.ErrorUserIDInvalid.Error())
		return uuid.Nil, errors.New(response.ErrorUserIDInvalid.Error())
	}
	userID, ok := id.(string)
	if !ok {
		loggers.Error(response.ErrorUserIDInvalid.Error())
		return uuid.Nil, errors.New(response.ErrorUserIDInvalid.Error())
	}
	StaffId, err := uuid.Parse(userID)
	if err != nil {
		loggers.Error(err)
		return uuid.Nil, err
	}
	return StaffId, nil
}
func GetUserRole(ctx *gin.Context) (string, error) {
	id, ok := ctx.Get(userRoleCtx)
	if !ok {
		return "", errors.New("user role not found")
	}
	idInt, ok := id.(string)
	if !ok {
		return "", errors.New("user role not found")
	}
	return idInt, nil
}
func GetPageQuery(ctx *gin.Context) (offset int64,
	err error) {
	offsetStr := ctx.DefaultQuery(page, cfg.APIProperty.DefaultPage)
	offset, err = strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		return 0, response.ErrorNotANumberPage
	}
	if offset < 0 {
		return 0, response.ErrorOffsetNotAUnsignedInt
	}
	return offset, nil
}
func GetPageSizeQuery(ctx *gin.Context) (limit int64,
	err error) {
	limitStr := ctx.DefaultQuery(pageSize, cfg.APIProperty.DefaultPageSize)
	limit, err = strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		return 0, response.ErrorNotANumberPageSize
	}
	if limit < 0 {
		return 0, response.ErrorLimitNotAUnsignedInt
	}
	return limit, nil
}
func CalculationPagination(page, pageSize int64) (offset, limit int64) {
	if page < 0 {
		page = 1
	}
	offset = (page - 1) * pageSize
	limit = pageSize
	return offset, limit
}
func ListPagination(ctx *gin.Context) (pagination model.Pagination, err error) {
	page, err := GetPageQuery(ctx)
	if err != nil {
		loggers.Error(err)
		return pagination, err
	}
	pageSize, err := GetPageSizeQuery(ctx)
	if err != nil {
		loggers.Error(err)
		return pagination, err
	}
	offset, limit := CalculationPagination(page, pageSize)
	pagination.Limit = limit
	pagination.Offset = offset
	pagination.Page = page
	pagination.PageSize = pageSize
	return pagination, nil
}
func GetUUIDParam(ctx *gin.Context, query string) (uuid.UUID, error) {
	queryData := ctx.Param(query)
	if queryData == "" {
		err := fmt.Sprintf(queryInvalid, queryData)
		return uuid.Nil, errors.New(err)
	}
	queryUUID, err := uuid.Parse(queryData)
	if err != nil {
		err := fmt.Sprintf(queryInvalid, queryData)
		return uuid.Nil, errors.New(err)
	}
	return queryUUID, nil
}
func CheckTime(query string) (time.Time, error) {
	if query != "" {
		checkTime, err := time.Parse(ParseDate, query)
		if err != nil {
			err := fmt.Sprintf(queryInvalid, ParseDate)
			return time.Time{}, errors.New(err)
		}
		return checkTime, nil
	}
	return time.Time{}, nil
}
func GetNullDateQuery(ctx *gin.Context, query string) (time.Time, error) {
	queryDate := GetNullStringQuery(ctx, query)
	date, err := CheckTime(queryDate)
	if err != nil {
		return time.Time{}, err
	}
	return date, nil
}
func GetNullUUIDQuery(ctx *gin.Context, query string) (uuid.UUID, error) {
	queryData := ctx.Query(query)
	if queryData != "" {
		queryUUID, err := uuid.Parse(queryData)
		if err != nil {
			err := fmt.Sprintf(queryInvalid, queryData)
			return uuid.Nil, errors.New(err)
		}
		return queryUUID, nil
	}
	return uuid.Nil, nil
}
func GetNullUUIDParam(ctx *gin.Context, query string) (uuid.UUID, error) {
	queryData := ctx.Param(query)
	if queryData != "" {
		queryUUID, err := uuid.Parse(queryData)
		if err != nil {
			err := fmt.Sprintf(queryInvalid, queryData)
			return uuid.Nil, errors.New(err)
		}
		return queryUUID, nil
	}
	return uuid.Nil, nil
}
func GetNullBooleanQuery(ctx *gin.Context, query string) (bool, error) {
	param := ctx.Query(query)
	if param != "" {
		boolVal, err := strconv.ParseBool(param)
		if err != nil {
			err := fmt.Sprintf(paramInvalid, query)
			return false, errors.New(err)
		}
		return boolVal, nil
	}
	return true, nil
}
func GetNullInt64Param(ctx *gin.Context, query string) (int64, error) {
	queryData := ctx.Query(query)
	if queryData != "" {
		queryInt, err := strconv.ParseInt(queryData, 10, 64)
		if err != nil {
			err := fmt.Sprintf(queryInvalid, queryData)
			return 0, errors.New(err)
		}
		return queryInt, nil
	}
	return 0, nil
}
func GetNullIntParam(ctx *gin.Context, query string) (int, error) {
	queryData := ctx.Query(query)
	if queryData != "" {
		queryInt, err := strconv.Atoi(queryData)
		if err != nil {
			err := fmt.Sprintf(queryInvalid, queryData)
			return 0, errors.New(err)
		}
		return queryInt, nil
	}
	return 0, nil
}
func GetNullFloat64Param(ctx *gin.Context, query string) (float64, error) {
	queryData := ctx.Query(query)
	if queryData != "" {
		queryFloat, err := strconv.ParseFloat(queryData, 64)
		if err != nil {
			err := fmt.Sprintf(queryInvalid, queryData)
			return 0, errors.New(err)
		}
		return queryFloat, nil
	}
	return 0, nil
}
func GetNullArrayStringQuery(ctx *gin.Context, query string) ([]string, error) {
	queryData := ctx.Query(query)
	if queryData != "" {
		chunks := strings.Split(queryData, ",")
		return chunks, nil
	}
	return []string{}, nil
}
func ResponseHeaderXTotalCountWrite(ctx *gin.Context, total int64) {
	ctx.Writer.Header().Set("X-Total-Count", strconv.Itoa(int(total)))
}
func FileTransfer(ctx *gin.Context, filePath string, contentType string) (err error) {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	ctx.Header("Content-Description", "File Transfer")
	ctx.Header("Content-Disposition", "attachment; filename="+path.Base(filePath))
	ctx.Data(http.StatusOK, "application/octet-stream", bytes)
	ctx.Writer.Header().Set("Content-Type", contentType)
	return nil
}
