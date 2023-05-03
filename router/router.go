package router

import (
	"github/Wuhao-9/go-gin-example/middleware"
	"github/Wuhao-9/go-gin-example/pkg/setting"
	v1 "github/Wuhao-9/go-gin-example/router/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(setting.RunMode)

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	router.GET("/auth", v1.GetAuth)

	apiv1 := router.Group("/api/v1")
	apiv1.Use(middleware.JWT())
	{
		// 获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		// 新建标签
		apiv1.POST("/tag", v1.AddTag)
		// 更新指定标签
		apiv1.PUT("/tag/:id", v1.EditTag)
		// 删除指定标签
		apiv1.DELETE("/tag/:id", v1.DeleteTag)

		// 获取指定文章
		apiv1.GET("/article/:id", v1.GetArticle)
		// 新增文章
		apiv1.POST("/article", v1.AddArticle)
		// 获取单页的所有文章
		apiv1.GET("/articles", v1.GetArticles)
		// 删除文章
		apiv1.DELETE("/article/:id", v1.DeleteArticle)
	}

	return router
}
