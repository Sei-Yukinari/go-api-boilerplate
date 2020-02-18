package router

import (
	"api/interfaces"
	"api/interfaces/user"
	"api/usecases"

	"github.com/gin-gonic/gin"
)

func Dispatch(r *gin.Engine, sqlHandler interfaces.SQLHandler, logger usecases.Logger) {
	apiV1 := r.Group("/api/v1")
	userRouter(apiV1, sqlHandler, logger)
}

func userRouter(r *gin.RouterGroup, sqlHandler interfaces.SQLHandler, logger usecases.Logger) {
	userController := user.NewUserController(sqlHandler, logger)
	user := r.Group("/user")
	user.GET("/", func(c *gin.Context) { userController.Index(c) })
	user.GET("/:id", func(c *gin.Context) { userController.Show(c) })
}
