package dao

import (
	"errors"
	"github/Wuhao-9/go-gin-example/models"
	"log"

	"gorm.io/gorm"
)

func GetTags(offset int, limit int, cst interface{}) (tags []models.Tag) {
	models.DB.Where(cst).Offset(offset).Limit(limit).Find(&tags)
	return
}

func GetTagTotal(cst interface{}) (count int64) {
	models.DB.Model(&models.Tag{}).Where(cst).Count(&count)
	return
}

func IsExist_tag_by_name(name string) bool {
	err := models.DB.Where("name = ?", name).First(&models.Tag{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false
		} else {
			log.Println(err)
			return false
		}
	}
	return true
}

func AddTag(name string, state int, createdBy string) bool {
	db := models.DB.Create(&models.Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})
	if db.Error != nil {
		return false
	}
	return true
}

func IsExist_tag_by_id(id int) bool {
	err := models.DB.Where(&models.Tag{Model: models.Model{ID: id}}).First(&models.Tag{}).Error
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return false
		} else {
			log.Println(err)
			return false
		}
	}
	return true
}

func EditTag(id int, data interface{}) bool {
	db := models.DB.Model(&models.Tag{}).Where("id = ?", id).Updates(data)
	if db.Error != nil {
		return false
	}
	return true
}

func DeleteTag(id int) bool {
	err := models.DB.Model(&models.Tag{}).Where("id = ?", id).Delete(models.Tag{}).Error
	if err != nil {
		return false
	}
	return true
}
