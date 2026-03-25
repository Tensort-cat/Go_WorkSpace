package route

import (
	"CurrencyExchangeApp/controller"
	"CurrencyExchangeApp/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoute() *gin.Engine {
	r := gin.Default()

	/*
		只允许 http://localhost:5173 这个前端，用 GET/POST/OPTIONS 方法跨域访问我，
		并且允许带 Origin、Content-Type、Authorization 请求头；前端只能读取 Content-Length 响应头；
		浏览器可以带凭证；预检结果缓存 12 小时
	*/
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	auth := r.Group("/api/auth")
	{
		auth.POST("/login", controller.Login)
		auth.POST("/register", controller.Register)
	}

	api := r.Group("/api")
	api.GET("/exchangeRates", controller.GetExchangeRates)
	api.Use(middleware.AuthMiddleware())
	{
		api.POST("/exchangeRates", controller.CreateExchangeRate)
	}

	article := r.Group("/api/articles")
	article.GET("", controller.GetArticle)
	article.GET("/", controller.GetArticle)
	article.GET("/:id", controller.GetArticleById)
	article.GET("/:id/like", controller.GetArticleLikes)
	article.Use(middleware.AuthMiddleware())
	{
		article.POST("", controller.CreateArticle)
		article.POST("/", controller.CreateArticle)
		article.POST("/:id/like", controller.LikeArticle)
	}

	return r
}
