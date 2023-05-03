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

var DB *gorm.DB

var (
	dbType, dbName, user, password, host, tablePrefix string
)

// define base model
type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

// define callback for update time
func (m *Model) BeforeCreate(tx *gorm.DB) error {
	m.CreatedOn = int(time.Now().Unix())
	m.ModifiedOn = int(time.Now().Unix())
	return nil
}

func (m *Model) BeforeUpdate(tx *gorm.DB) error {
	m.ModifiedOn = int(time.Now().Unix())
	return nil
}

func init() {
	var err error
	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatalf("Fail to get section 'database': %v", err)
	}

	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("DBNAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PWD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true&loc=Local", user, password, host, dbName)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB.NamingStrategy = schema.NamingStrategy{ // 设置表明和列名命名策略
		TablePrefix:   tablePrefix,
		SingularTable: true,
		NameReplacer:  nil,
		NoLowerCase:   false,
	}

	// err = DB.AutoMigrate(&Tag{}) // 自动迁移模型
	err = DB.AutoMigrate(&Article{})
	if err != nil {
		fmt.Println(err)
	}

	err = DB.AutoMigrate(&Auth{})
	if err != nil {
		fmt.Println(err)
	}
}
