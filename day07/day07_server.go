package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserProfile struct {
	Username string `json:"username"`
	Age      int    `json:"age"`
	Status   string `json:"status"`
}

func main071() {
	// 创建一个默认的路由引擎
	// 想象它是一个“前台经理”，负责分配客人的请求
	r := gin.Default()

	// 定义一个“路由”(Route)
	// 意思是： 如果有人访问 http://locathost:8080/hello
	r.GET("/hello", func(c *gin.Context) {
		// 回复一个JSON数据
		c.JSON(http.StatusOK, gin.H{
			"message": "你好，Gopher！你的第一个Web服务器已上线!",
		})
	})

	// 4. 定义一个获取用户信息的接口
	r.GET("/user", func(c *gin.Context) {
		user := UserProfile{
			Username: "二次元高手",
			Age:      18,
			Status:   "在线",
		}
		c.JSON(http.StatusOK, user)
	})

	// 魔改任务：添加一个路由/ping， 返回{"result":"pong"}
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"result": "pong",
		})
	})

	// 启动服务器，默认监听 8080 端口
	fmt.Println("服务器已启动，监听 http://localhost:9000")
	r.Run(":9000") // 修改端口号为9000
}
