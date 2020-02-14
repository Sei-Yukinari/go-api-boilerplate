package http

import (
	"fmt"

	"api/config"

	"github.com/gin-gonic/gin"
)

func StartHttpServer() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hoge",
		})
	})
	c := config.Load()
	if err := r.Run(":" + c.Server.Port); err != nil {
		fmt.Printf("Error : [%s]", err)
	}

}
