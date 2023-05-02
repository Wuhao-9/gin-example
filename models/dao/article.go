package dao

import (
	"fmt"
	"github/Wuhao-9/go-gin-example/models"
)

func AddArticle(record_info *models.Article) bool {
	err := models.DB.Model(&models.Article{}).Create(record_info).Error
	if err != nil {
		return false
	}
	return true
}

func IsExist_article_by_id(id int) bool {
	err := models.DB.Where("id = ?", id).First(&models.Article{}).Error
	if err != nil {
		return false
	}
	return true
}

func GetArticle(id int) (article models.Article) {
	err := models.DB.Preload("Tag").Where("id = ?", id).First(&article).Error
	if err != nil {
		fmt.Println(err)
	}

	return
}

func GetArticleTotal(cst interface{}) (count int64) {
	models.DB.Model(&models.Article{}).Where(cst).Count(&count)
	return
}

func GetArticles(offset int, limit int, cst interface{}) (articles []models.Article, ok bool) {
	err := models.DB.Preload("Tag").Where(cst).Offset(offset).Limit(limit).Find(&articles).Error
	if err != nil {
		ok = false
	} else {
		ok = true
	}
	return
}

func DeleteArticle(target int) {
	models.DB.Delete(models.Article{}, target)
}
