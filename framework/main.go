package main

import (
	"framework/handler"

	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

func main() {
	ginDemo()
}

func ginDemo() {
	var r = gin.Default() // 初始化 Gin 引擎，带 Logger + Recovery 中间件
	//1.请求进来，执行recover
	//2.程序异常，抛出panic
	//3.panic被 recover捕获，返回异常信息，并Abort,终止这次请求
	r.Use(handler.Recover)
	r.GET("/ping", ping)
	var err = r.Run(":8084")
	if err != nil {
		return
	} // 启动服务，监听 8080 端口
}

func ping(c *gin.Context) {
	var query = c.DefaultQuery("a", "b")
	c.JSON(200, gin.H{
		"message": query,
	})
	var slice = []int{1, 2, 3, 4, 5}
	slice[6] = 6
	c.JSON(200, gin.H{
		"message": "ping",
	})
	log.Info("执行到")
}
