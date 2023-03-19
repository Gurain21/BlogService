package routers

import (
	"github.com/gin-gonic/gin"
	"jaingke2023.com/BlogService/middleware"
	"jaingke2023.com/BlogService/routers/api"
	"jaingke2023.com/BlogService/routers/api/v1"
)

// InitRouterGroups 初始化所有的路由组
func InitRouterGroups() *gin.Engine {
	
	router := gin.New()
	//使用中间件，开启gin的Logger日志和recover拦截 中间件
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	
	//  /auth  用户 的 token 令牌 发放接口
	router.GET("/auth", api.CheckAuth)
	
	//  接入自己写的 JWT 验证中间件 ，注意位置！用户登录不能拦截了！！！
	
	router.Use(middleware.JWT())
	
	apiV1 := router.Group("/api/v1")
	
	{
		//获取标签列表
		apiV1.GET("/tags", v1.GetTags)
		//新建标签
		apiV1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiV1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiV1.DELETE("/tags/:id", v1.DeleteTag)
		
		//获取文章列表
		apiV1.GET("/articles", v1.GetArticles)
		//新建文章
		apiV1.POST("/articles/", v1.AddArticle)
		//获取指定的文章
		apiV1.POST("/articles/:id", v1.GetArticle)
		//更新指定的文章
		apiV1.PUT("/articles/:id", v1.EditArticle)
		//删除指定的文章
		apiV1.DELETE("/articles/:id", v1.DeleteArticle)
		
	}
	
	return router
}
