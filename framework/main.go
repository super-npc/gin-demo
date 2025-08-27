package main

import "github.com/gin-gonic/gin"

func main() {
	ginDemo()
}

func ginDemo() {
	var r = gin.Default() // 初始化 Gin 引擎，带 Logger + Recovery 中间件
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
	c.JSON(200, gin.H{
		"message": "ping",
	})
}
