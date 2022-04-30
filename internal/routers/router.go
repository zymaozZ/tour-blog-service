package routers

import (
	v1 "blog-service/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	article := v1.NewArticle()
	tag := v1.NewTag()

	apiV1 := r.Group("/api/v1")
	{
		// 创建标签
		apiV1.POST("/tags", tag.Create)
		// 删除指定标签
		apiV1.DELETE("/tags/:id", tag.Delete)
		// 更新指定标签
		apiV1.PUT("/tags/:id", tag.Update)
		// 获取标签列表
		apiV1.GET("/tags", tag.List)

		// 创建文章
		apiV1.POST("/articles", article.Create)
		// 删除指定文章
		apiV1.DELETE("/articles/:id", article.Delete)
		// 更新指定文章
		apiV1.PUT("/articles/:id", article.Update)
		// 获取指定文章
		apiV1.GET("/articles/:id", article.Get)
		// 获取文章列表
		apiV1.GET("/articles", article.List)
	}

	return r
}
