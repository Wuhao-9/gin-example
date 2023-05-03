package dao

import (
	"github/Wuhao-9/go-gin-example/models"
	"log"
)

func CheckAuth(account, passwd string) bool {
	err := models.DB.Where(models.Auth{Account: account, Passwd: passwd}).First(&models.Auth{}).Error
	if err != nil {
		log.Printf("CheckAuth fail: %v", err)
		return false
	}
	return true
}
