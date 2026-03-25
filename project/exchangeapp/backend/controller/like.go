package controller

import (
	"CurrencyExchangeApp/dao"
	"CurrencyExchangeApp/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

func LikeArticle(ctx *gin.Context) {
	id := ctx.Param("id")
	redisKey := buildRedisKey(id)

	_, err := dao.RedisTemplate.Get(redisKey).Result()
	switch err {
	case nil: // redis 有
		// 清除 redis 缓存
		dao.RedisTemplate.Del(redisKey)
	case redis.Nil: // redis 没有
		break
	default: // 出错了
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// 更新 mysql 数据
	dao.DB.Model(&model.Article{}).
		Where("id = ?", id).
		UpdateColumn("likes", gorm.Expr("likes + ?", 1))

	ctx.JSON(http.StatusOK, gin.H{"msg": "success"})
}

func GetArticleLikes(ctx *gin.Context) {
	id := ctx.Param("id")
	redisKey := buildRedisKey(id)

	likes, err := dao.RedisTemplate.Get(redisKey).Result()
	if err == redis.Nil { // Nil reply Redis returns when key does not exist.
		article, err := model.GetArticleById(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		likes = strconv.Itoa(article.Likes)
		dao.RedisTemplate.Set(redisKey, likes, 100*time.Second)
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"likes": likes})
}

func buildRedisKey(id string) string {
	return "article:" + id + ":likes"
}
