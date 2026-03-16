package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")
	r.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	// 单个文件
	// r.POST("/upload", func(ctx *gin.Context) {
	// 	// 从请求中读取文件
	// 	f, err := ctx.FormFile("f1")
	// 	if err != nil {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// 	// 将文件保存到服务器
	// 	dst := fmt.Sprintf("./%s", f.Filename)
	// 	ctx.SaveUploadedFile(f, dst)
	// })

	// 多个文件
	r.POST("/uploads", func(ctx *gin.Context) {
		// 从请求中读取多个文件
		form, err := ctx.MultipartForm()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		files := form.File["f1"]
		for _, f := range files {
			dst := fmt.Sprintf("./%s", f.Filename)
			ctx.SaveUploadedFile(f, dst)
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "上传成功"})
	})

	r.Run()

}
