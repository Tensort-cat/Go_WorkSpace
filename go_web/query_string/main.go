package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/query", func(c *gin.Context) {
		// name := c.Query("name")
		// name := c.DefaultQuery("name", "匿名用户") // 没输入name参数时，默认值为匿名用户
		name, ok := c.GetQuery("name") // GetQuery返回两个值，第一个是参数值，第二个是参数是否存在
		if !ok {
			name = "匿名用户"
		}

		// age := c.Query("age")
		// age := c.DefaultQuery("age", "未知") // 没输入age参数时，默认值为未知
		age, ok := c.GetQuery("age") // GetQuery返回两个值，第一个是参数值，第二个是参数是否存在
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
