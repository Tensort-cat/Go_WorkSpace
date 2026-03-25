package model

import (
	"CurrencyExchangeApp/dao"
	"log"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
}

func CreateUser(user *User) (err error) {
	if err = dao.DB.AutoMigrate(&user); err != nil {
		log.Printf("Create user failed: %v", err)
		return
	}

	if err = dao.DB.Create(&user).Error; err != nil {
		log.Printf("Create user failed: %v", err)
		return
	}
	return
}
