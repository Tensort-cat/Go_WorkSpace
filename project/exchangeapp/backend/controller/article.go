package controller

import (
	"CurrencyExchangeApp/dao"
	"CurrencyExchangeApp/model"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

const (
	CACHE_KEY = "article"
)

func CreateArticle(ctx *gin.Context) {
	var article model.Article
	if err := ctx.ShouldBindJSON(&article); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := model.CreateArticle(&article); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 创建成功时，删除 redis 中的旧缓存，并向前端返回数据
	dao.RedisTemplate.Del(CACHE_KEY)
	ctx.JSON(http.StatusCreated, article)
}

func GetArticle(ctx *gin.Context) {
	cacheData, err := dao.RedisTemplate.Get(CACHE_KEY).Result()
	if err == redis.Nil {
		var articles []model.Article
		if err := model.GetArticles(&articles); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// 先返回主数据
		ctx.JSON(http.StatusOK, articles)

		// 再尝试写缓存；失败不影响本次请求
		if toJSON, err := json.Marshal(articles); err == nil {
			_ = dao.RedisTemplate.Set(CACHE_KEY, toJSON, 5*time.Minute).Err()
		}
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Data(http.StatusOK, "application/json; charset=utf-8", []byte(cacheData))
}

func GetArticleById(ctx *gin.Context) {
	id := ctx.Param("id")
	article, err := model.GetArticleById(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, article)
}
