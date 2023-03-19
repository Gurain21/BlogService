package util

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"jaingke2023.com/BlogService/pkg/settings"
)

//GetPage 分页查找
func GetPage(c *gin.Context) (result int) {
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * settings.PageSize
		
	}
	return
}
