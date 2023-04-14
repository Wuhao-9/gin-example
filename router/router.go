package router

import (
	"github/Wuhao-9/go-gin-example/pkg/setting"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(setting.RunMode)
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	return router
}
