package api

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"jaingke2023.com/BlogService/models"
	"jaingke2023.com/BlogService/pkg/e"
	"jaingke2023.com/BlogService/pkg/logging"
	"jaingke2023.com/BlogService/pkg/util"
	"net/http"
)

func CheckAuth(c *gin.Context) {
	valid := validation.Validation{}
	code := e.SUCCESS
	data := make(map[string]interface{})
	
	username := c.Query("username")
	password := c.Query("password")
	
	valid.Required(username, "username").Message("The argument username is required")
	valid.MaxSize(username, 50, "username").Message("The argument username max size is 50")
	valid.Required(password, "password").Message("The argument password is required")
	valid.MaxSize(password, 50, "password").Message("The argument password max size is 50")
	
	//a := models.Auth{Username: username, Password: password}
	//ok, _ := valid.Valid(&a)
	
	if !valid.HasErrors() {
		isExist := models.CheckAuth(username, password)
		
		if isExist {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = e.ErrorAuthToken
			}
			data["token"] = token
		} else {
			code = e.ErrorAuth
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
		"data": data,
	})
	
}
