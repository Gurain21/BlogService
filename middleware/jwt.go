package middleware

import (
	"github.com/gin-gonic/gin"
	"jaingke2023.com/BlogService/pkg/e"
	"jaingke2023.com/BlogService/pkg/util"
	"net/http"
	"time"
)

func JWT2() gin.HandlerFunc {
	return func(c *gin.Context) {
		
		var code int
		var data interface{}
		
		code = e.SUCCESS
		token := c.Query("token")
		if token == "" {
			code = e.InvalidParams
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ErrorAuthCheckTokenFail
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ErrorAuthCheckTokenTimeout
			}
		}
		
		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})
			//终止其相关的请求处理
			c.Abort()
			return
		}
		//放行
		c.Next()
	}
}

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code = e.SUCCESS
		var data = make(map[string]interface{})
		
		token := c.Query("token")
		
		if token == "" {
			code = e.InvalidParams
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ErrorAuthCheckTokenFail
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ErrorAuthCheckTokenTimeout
			}
		}
		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})
			//终止其相关的请求处理
			c.Abort()
			return
		}
		
		c.Next()
		
	}
}
