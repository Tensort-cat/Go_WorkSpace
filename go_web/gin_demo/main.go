package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建带默认中间件（日志与恢复）的 Gin 路由器
	r := gin.Default()

	// 定义简单的 GET 路由
	r.GET("/ping", func(c *gin.Context) {
		// 返回 JSON 响应
		c.JSON(http.StatusOK, gin.H{
			"msg":  "pong",
			"data": "GET",
		})
	})

	// 定义简单的 POST 路由
	r.POST("/ping", func(c *gin.Context) {
		// 返回 JSON 响应
		c.JSON(http.StatusOK, gin.H{
			"msg":  "pong",
			"data": "POST",
		})
	})

	// 定义简单的 PUT 路由
	r.PUT("/ping", func(c *gin.Context) {
		// 返回 JSON 响应
		c.JSON(http.StatusOK, gin.H{
			"msg":  "pong",
			"data": "PUT",
		})
	})

	// 定义简单的 DELETE 路由
	r.DELETE("/ping", func(c *gin.Context) {
		// 返回 JSON 响应
		c.JSON(http.StatusOK, gin.H{
			"msg":  "pong",
			"data": "DELETE",
		})
	})

	// 定义一个返回结构体的 GET 路由
	r.GET("/struct", func(ctx *gin.Context) {
		var data struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}
		data.Name = "张三"
		data.Age = 18

		ctx.JSON(http.StatusOK, data)
	})

	// 定义一个同时处理多种 HTTP 方法的路由
	r.Any("/any", func(ctx *gin.Context) {
		switch ctx.Request.Method {
		case http.MethodGet:
			ctx.JSON(http.StatusOK, gin.H{"msg": "GET request"})
		case http.MethodPost:
			ctx.JSON(http.StatusOK, gin.H{"msg": "POST request"})
		case http.MethodPut:
			ctx.JSON(http.StatusOK, gin.H{"msg": "PUT request"})
		case http.MethodDelete:
			ctx.JSON(http.StatusOK, gin.H{"msg": "DELETE request"})
		default:
			ctx.JSON(http.StatusOK, gin.H{"msg": "Other request"})
		}
	})

	// 默认端口 8080 启动服务器
	// 监听 0.0.0.0:8080（Windows 下为 localhost:8080）
	r.Run()
}
