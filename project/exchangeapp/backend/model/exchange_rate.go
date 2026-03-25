package model

import (
	"CurrencyExchangeApp/dao"
	"time"
)

type ExchangeRate struct {
	ID           uint      `gorm:"primarykey" json:"id"`
	FromCurrency string    `json:"fromCurrency" binding:"required"`
	ToCurrency   string    `json:"toCurrency" binding:"required"`
	Rate         float64   `json:"rate" binding:"required"`
	Date         time.Time `json:"date"`
}

func CreateExchangeRate(ecr *ExchangeRate) error {
	if err := dao.DB.AutoMigrate(&ExchangeRate{}); err != nil {
		return err
	}

	return dao.DB.Create(ecr).Error
}

func GetExchangeRates(ecrs *[]ExchangeRate) error {
	res := dao.DB.Find(ecrs)
	return res.Error
}
