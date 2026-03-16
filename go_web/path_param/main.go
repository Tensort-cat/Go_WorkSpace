package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `uri:"name" json:"name"`
	Age  string `uri:"age" json:"age"`
}

func main() {
	r := gin.Default()

	r.GET("path_param/:name/:age", func(ctx *gin.Context) {
		// ctx.JSON(http.StatusOK, gin.H{
		// 	"name": ctx.Param("name"),
		// 	"age":  ctx.Param("age"),
		// })

		// 参数绑定
		var user User
		if err := ctx.ShouldBindUri(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, user)

	})

	r.POST("/json", func(ctx *gin.Context) {
		var user User
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Printf("%v", user)
		ctx.JSON(http.StatusOK, user)
	})

	r.Run()
}
