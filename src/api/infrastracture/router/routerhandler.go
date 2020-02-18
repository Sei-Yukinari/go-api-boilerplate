package router

import (
	"api/interfaces"
	"api/usecases"

	"github.com/gin-gonic/gin"
)

func Handler(sqlHandler interfaces.SQLHandler, logger usecases.Logger) *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	Dispatch(r, sqlHandler, logger)
	return r
}
