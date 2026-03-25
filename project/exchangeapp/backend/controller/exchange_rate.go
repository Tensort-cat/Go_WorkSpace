package controller

import (
	"CurrencyExchangeApp/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateExchangeRate(ctx *gin.Context) {
	var ecr model.ExchangeRate
	if err := ctx.ShouldBindJSON(&ecr); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ecr.Date = time.Now()
	if err := model.CreateExchangeRate(&ecr); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, ecr)
}

func GetExchangeRates(ctx *gin.Context) {
	var ecrs []model.ExchangeRate
	err := model.GetExchangeRates(&ecrs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, ecrs)
}
