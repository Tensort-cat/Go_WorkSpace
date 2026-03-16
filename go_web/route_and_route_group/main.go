package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"msg": "不存在的页面",
		})
	})

	// 路由组
	videoGroup := r.Group("/video")
	{
		videoGroup.GET("/index", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "视频首页",
			})
		})
		videoGroup.GET("/detail", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "视频详情页",
			})
		})

		// 路由组支持嵌套
		xx := videoGroup.Group("/xx")
		// 访问http://localhost:8080/video/xx/xx
		xx.GET("/xx", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, "xx") })
	}

	r.Run()
}
