package minio

import (
	"EduCRM/package/service"
	"EduCRM/tools/logger"

	"github.com/gin-gonic/gin"
)

type MinIOHandler struct {
	MinIOEndPoint
}
type MinIOEndPoint interface {
	UploadImages(ctx *gin.Context)
	UploadImage(ctx *gin.Context)
	UploadDoc(ctx *gin.Context)
	DownloadFile(ctx *gin.Context)
	DownloadAssetFile(ctx *gin.Context)
}

func NewMinIOHandler(service *service.Service,
	loggers *logger.Logger) *MinIOHandler {
	return &MinIOHandler{
		MinIOEndPoint: NewMinIOEndPointHandler(service, loggers),
	}
}
