package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func indexHandler(ctx *gin.Context) {
	fmt.Println("我是 indexHandler")
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

// 定义一个中间件 m1
func m1(ctx *gin.Context) {
	fmt.Println("我是中间件 m1")
	start := time.Now()
	ctx.Next() // 调用后续的处理函数
	cost := time.Since(start)
	fmt.Printf("m1 执行完成，总耗时 %v\n", cost)
}

// 定义一个中间件 m2
func m2(ctx *gin.Context) {
	fmt.Println("我是中间件 m2")
	ctx.Abort() // 阻止后续的处理函数被调用
	fmt.Println("m2 执行完成")
}

// 定义一个检查用户登录的中间件
func authMiddleware() gin.HandlerFunc {
	// 这里一般会放一些准备工作，比如连接数据库、读取配置等
	// ...
	return func(ctx *gin.Context) {
		// 伪代码
		/*
			if !isLoggedIn(ctx) {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "请先登录"})
				ctx.Abort()
				return
			}
			ctx.Next()
		*/
	}
}

func main() {
	r := gin.Default()

	// r.GET("/index", m1, indexHandler)

	// 全局中间件
	r.Use(m1, m2)

	// 均不会执行 indexHandler
	r.GET("/index", indexHandler)
	r.POST("/index", indexHandler)
	r.PUT("/index", indexHandler)

	r.Run()
}
