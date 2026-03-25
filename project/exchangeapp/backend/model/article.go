package model

import (
	"CurrencyExchangeApp/dao"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title   string `binding:"required"` // 文章标题
	Content string `binding:"required"` // 文章内容
	Preview string `binding:"required"` // 预览
	Likes   int    `gorm:"default:0"`   // 点赞数
}

func CreateArticle(article *Article) error {
	err := dao.DB.AutoMigrate(&Article{})
	if err != nil {
		return err
	}

	return dao.DB.Create(article).Error
}

func GetArticles(articles *[]Article) error {
	res := dao.DB.Find(&articles)
	return res.Error
}

func GetArticleById(id any) (Article, error) {
	var article Article
	res := dao.DB.First(&article, id)
	if err := res.Error; err != nil {
		return Article{}, err
	}
	return article, nil
}
