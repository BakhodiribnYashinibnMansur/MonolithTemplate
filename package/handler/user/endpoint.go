package user

import (
	"EduCRM/api/middleware"

	"github.com/gin-gonic/gin"
)

func UserRouter(api *gin.Engine, handler *UserHandler) {
	user := api.Group("/api/v1/user")
	{
		user.POST("/sign-in", handler.UserEndPoint.SignInUser)
	}
	userAuth := api.Group("/api/v1/user", middleware.AuthRequestHandler)
	{
		userAuth.POST("/create", handler.UserEndPoint.CreateUser)
		userAuth.PUT("/update/:id", handler.UserEndPoint.UpdateUser)
		userAuth.DELETE("/delete/:id", handler.UserEndPoint.DeleteUser)
		userAuth.GET("/list", handler.UserEndPoint.GetUserList)
		userAuth.GET("/:id", handler.UserEndPoint.GetUserByID)
		userAuth.GET("/me", handler.UserEndPoint.GetUserMe)
		userAuth.PUT("/update-password/:id", handler.UserEndPoint.UpdateUserPassword)
	}
}
