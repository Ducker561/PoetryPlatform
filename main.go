package main

import (
	"github.com/gin-gonic/gin"
	"InternetPlus/conf"
	"InternetPlus/routes"

	"fmt"
)

// @title InternetPlus API
// @version 0.0.1
// @description This is a sample Server pets
// @name Alesp
// @BasePath /api/v1
func main() { // http://localhost:3000/swagger/index.html
	//从配置文件读入配置
	gin.SetMode(gin.DebugMode)
	conf.Init()
	//转载路由 swag init -g main.go
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)
	fmt.Println("初始化完成")
}
