package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/index", func(ctx *gin.Context) {
		// ctx.JSON(http.StatusOK, gin.H{
		// 	"status": "ok",
		// })
		ctx.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})

	r.GET("/a", func(ctx *gin.Context) {
		// 跳转到 /b 对应的处理函数
		ctx.Request.URL.Path = "/b" // 修改请求的 URL 路径
		r.HandleContext(ctx)        // 将请求交给路由器处理
	})

	r.GET("/b", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	r.Run()
}
