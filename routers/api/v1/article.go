package v1

import (
	"jaingke2023.com/BlogService/pkg/logging"
	"net/http"
	
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	
	"jaingke2023.com/BlogService/models"
	"jaingke2023.com/BlogService/pkg/e"
	"jaingke2023.com/BlogService/pkg/settings"
	"jaingke2023.com/BlogService/pkg/util"
)

// GetArticle 获取单个文章
func GetArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	
	var data = make(map[string]interface{})
	
	valid := validation.Validation{}
	
	code := e.InvalidParams
	valid.Min(id, 1, "id").Message("The argument id min value is 1 ")
	
	if !valid.HasErrors() {
		data["article"] = models.GetArticleById(id)
		code = e.SUCCESS
	} else {
		for _, v := range valid.Errors {
			logging.Info(v.Error())
			
		}
	}
	
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
	
}

// GetArticles 获取多个文章
func GetArticles(c *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	valid := validation.Validation{}
	
	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
		
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}
	
	var tagId int = -1
	if arg := c.Query("tag_id"); arg != "" {
		tagId = com.StrTo(arg).MustInt()
		maps["tag_id"] = tagId
		
		valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	}
	
	code := e.InvalidParams
	if !valid.HasErrors() {
		code = e.SUCCESS
		
		data["lists"] = models.GetArticles(util.GetPage(c), settings.PageSize, maps)
		data["total"] = models.GetArticleTotal(maps)
		
	} else {
		for _, v := range valid.Errors {
			logging.Info(v.Error())
			
		}
	}
	
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// AddArticle 新增文章
func AddArticle(c *gin.Context) {
	id := com.StrTo(c.Query("tagId")).MustInt()
	state := com.StrTo(c.Query("state")).MustInt()
	
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	createBy := c.Query("createBy")
	
	models.AddArticle(id, state, title, desc, content, createBy)
}

// EditArticle 修改文章
func EditArticle(c *gin.Context) {
	valid := validation.Validation{}
	
	id := com.StrTo(c.Param("id")).MustInt()
	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	modifiedBy := c.Query("modified_by")
	
	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}
	
	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.MaxSize(title, 100, "title").Message("标题最长为100字符")
	valid.MaxSize(desc, 255, "desc").Message("简述最长为255字符")
	valid.MaxSize(content, 65535, "content").Message("内容最长为65535字符")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
	
	code := e.InvalidParams
	if !valid.HasErrors() {
		if models.ExistArticleById(id) {
			if models.ExistTagById(tagId) {
				data := make(map[string]interface{})
				if tagId > 0 {
					data["tag_id"] = tagId
				}
				if title != "" {
					data["title"] = title
				}
				if desc != "" {
					data["desc"] = desc
				}
				if content != "" {
					data["content"] = content
				}
				
				data["modified_by"] = modifiedBy
				
				models.UpdateArticleById(id, data)
				code = e.SUCCESS
			} else {
				code = e.ErrorNotExistTag
			}
		} else {
			code = e.ErrorNotExistArticle
		}
	} else {
		for _, v := range valid.Errors {
			logging.Info(v.Error())
			
		}
	}
	
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// DeleteArticle 删除文章
func DeleteArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("The argument id must large than 1")
	code := e.InvalidParams
	if !valid.HasErrors() {
		models.DeleteArticleById(id)
		code = e.SUCCESS
		
	} else {
		for _, v := range valid.Errors {
			logging.Info(v.Error())
			
		}
	}
	
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}
