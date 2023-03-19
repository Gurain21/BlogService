package main

import (
	"fmt"
	"jaingke2023.com/BlogService/pkg/logging"
	"net/http"
	"strconv"
	
	"github.com/gin-gonic/gin"
	
	"jaingke2023.com/BlogService/pkg/settings"
	"jaingke2023.com/BlogService/routers"
)

func main() {
	gin.SetMode(gin.DebugMode)
	
	//default := gin.Default()
	//default.GET("/", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"test": "测试数据",
	//		"msg":  "ok",
	//	})
	//})
	router := routers.InitRouterGroups()
	
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", settings.HttpPort),
		Handler:        router,
		ReadTimeout:    settings.ReadTimeOut,
		WriteTimeout:   settings.WriteTimeOut,
		MaxHeaderBytes: 1 << 20,
	}
	
	logging.Info("Start a very good web service and listen address: http://localhost:", strconv.Itoa(settings.HttpPort))
	err := server.ListenAndServe()
	if err != nil {
		logging.Error(err.Error())
	}
	
}
