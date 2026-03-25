package middleware

import (
	"CurrencyExchangeApp/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 用于检验登录请求头的中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "未授权"})
			ctx.Abort()
			return
		}

		// 解析JWT
		username, err := util.ParseJWT(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invaild token"})
			ctx.Abort()
			return
		}
		ctx.Set("username", username)
		ctx.Next()
	}
}
