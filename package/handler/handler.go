package handler

import (
	"EduCRM/api/middleware"
	"EduCRM/config"
	"EduCRM/package/handler/minio"
	"EduCRM/package/handler/user"
	"EduCRM/package/service"
	"EduCRM/tools/logger"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

type Handler struct {
	User  *user.UserHandler
	MinIO *minio.MinIOHandler
}

func NewHandler(service *service.Service,
	loggers *logger.Logger) *Handler {
	return &Handler{
		User:  user.NewUserHandler(service, loggers),
		MinIO: minio.NewMinIOHandler(service, loggers),
	}
}
func (handler *Handler) InitRoutes() (route *gin.Engine) {
	cfg := config.Config()
	route = gin.New()
	gin.SetMode(gin.ReleaseMode)
	if cfg.Server.Environment == "development" {
		gin.SetMode(gin.DebugMode)
	}
	route.HandleMethodNotAllowed = true
	middleware.GinMiddleware(route)
	//swagger settings
	docs.SwaggerInfo.Title = cfg.Server.AppName
	docs.SwaggerInfo.Version = cfg.Server.AppVersion
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	serverHost := ""
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler),
		func(ctx *gin.Context) {
			serverHost = ctx.Request.Host
		})
	docs.SwaggerInfo.Host = serverHost
	route.Static("/public", "./public/")
	//routers
	minio.MinIORouter(route, handler.MinIO)
	user.UserRouter(route, handler.User)
	return
}
