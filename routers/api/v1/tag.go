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

//GetTags  get a tags of article ,
func GetTags(c *gin.Context) {
	name := c.Query("name")
	
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	
	if name != "" {
		maps["name"] = name
	}
	
	var state int = -1
	
	arg := c.Query("state")
	if arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}
	
	code := e.SUCCESS
	
	data["lists"] = models.GetTags(util.GetPage(c), settings.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)
	
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
	
}

//AddTag  get a tags of article ,
func AddTag(c *gin.Context) {
	//1.获取requestURL中传递过来的参数
	//name=1&state=1&created_by=test
	name := c.Query("name")
	state := com.StrTo(c.Query("state")).MustInt()
	createdBy := c.Query("created_by")
	
	//2.参数校验，判断参数是否符合业务逻辑
	valid := validation.Validation{}
	valid.Required(name, "name").Message("The argument name is required!")
	valid.MaxSize(name, 60, "name").Message("The argument name's length is  out of range max size '60' ! ")
	
	valid.Required(createdBy, "created_by").Message("The argument created_by is required!")
	valid.MaxSize(createdBy, 60, "created_by").Message("The argument created_by's length is  out of range max size '60' ! ")
	
	valid.Range(state, 0, 1, "state").Message("The argument state is only in range 0 or 1")
	
	code := e.InvalidParams
	//3.
	if !valid.HasErrors() {
		if !models.ExistTagByName(name) {
			
			if models.AddTag(name, createdBy, state) {
				code = e.SUCCESS
			} else {
				code = e.ERROR
			}
			
		} else {
			code = e.ErrorExistTag
		}
	} else {
		code = e.InvalidParams
		
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

//EditTag  get a tags of article ,
func EditTag(c *gin.Context) {
	// restful 风格 采用 Param方法接收
	id := com.StrTo(c.Param("id")).MustInt()
	
	name := c.Query("name")
	modifiedBy := c.Query("modified_by")
	state := com.StrTo(c.Query("state")).MustInt()
	
	//	valid 校验参数
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("id must large than 1")
	valid.Required(name, "name").Message("The argument name is required!")
	valid.MaxSize(name, 60, "name").Message("The argument name's length is  out of range max size '60' ! ")
	
	valid.Required(modifiedBy, "modified_by").Message("The argument modified_by is required!")
	valid.MaxSize(modifiedBy, 60, "modified_by").Message("The argument modified_by's length is  out of range max size '60' ! ")
	valid.Range(state, 0, 1, "state").Message("The argument state is only in range 0 or 1")
	
	// 修改数据返回结果
	code := e.InvalidParams
	if !valid.HasErrors() {
		if !models.ExistTagById(id) {
			code = e.ErrorNotExistTag
		} else {
			maps := make(map[string]interface{})
			maps["name"] = name
			maps["modified_by"] = modifiedBy
			maps["state"] = state
			maps["id"] = id
			models.UpdateTag(maps)
			code = e.SUCCESS
		}
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

//DeleteTag  get a tags of article ,
func DeleteTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	
	valid := validation.Validation{}
	
	valid.Min(id, 1, "id")
	
	var code int
	//若id参数没有错误
	if !valid.HasErrors() {
		//	判断数据库中是否存在该条tag，有就删除，没有返回err
		if models.ExistTagById(id) {
			if models.DeleteTagById(id) {
				code = e.SUCCESS
			}
			
		} else {
			code = e.ErrorNotExistTag
		}
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
