package models

import (
	"fmt"
	"github/Wuhao-9/go-gin-example/pkg/setting"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

var (
	dbType, dbName, user, password, host, tablePrefix string
)

type Model struct {
	ID         int       `gorm:"primary_key" json:"id"`
	CreatedOn  time.Time `json:"created_on"`
	ModifiedOn time.Time `json:"modified_on"`
}

func init() {
	var err error
	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("DBNAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PWD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true&loc=Local", user, password, host, dbName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.NamingStrategy = schema.NamingStrategy{ // 设置表明和列名命名策略
		TablePrefix:   tablePrefix,
		SingularTable: true,
		NameReplacer:  nil,
		NoLowerCase:   false,
	}

	err = db.AutoMigrate(&Model{}) // 自动迁移模型
	if err != nil {
		log.Println(err)
	}
}
