package main

import (
	"Jankinwu-gin-vue/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", controller.Ping)
	r.GET("/api/user/register", controller.Register)
	r.Run()
}
