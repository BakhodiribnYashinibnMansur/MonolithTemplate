package minio

import (
	"EduCRM/api/middleware"

	"github.com/gin-gonic/gin"
)

func MinIORouter(api *gin.Engine, handler *MinIOHandler) {
	minio := api.Group("/api/v1/minio", middleware.AuthRequestHandler)
	{
		minio.POST("/upload-images", handler.MinIOEndPoint.UploadImages)
		minio.POST("/upload-image", handler.MinIOEndPoint.UploadImage)
		minio.POST("/upload-doc", handler.MinIOEndPoint.UploadDoc)
	}
	api.GET("/api/v1/download-file/:file-path", handler.MinIOEndPoint.DownloadFile)
	// api.GET("/public/", handler.MinIOEndPoint.DownloadAssetFile)
}
