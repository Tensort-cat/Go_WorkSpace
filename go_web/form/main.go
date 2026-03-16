package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/form", func(c *gin.Context) {
		// name := c.PostForm("name")
		// name := c.DefaultPostForm("name", "匿名用户") // 没输入name参数时，默认值为匿名用户
		name, ok := c.GetPostForm("name")
		if !ok {
			name = "匿名用户"
		}

		// age := c.PostForm("age")
		// age := c.DefaultPostForm("age", "未知") // 没输入age参数时，默认值为未知
		age, ok := c.GetPostForm("age")
		if !ok {
			age = "未知"
		}

		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})
	r.Run()
}
