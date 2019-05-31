package model

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/mumushuiding/go-attendance/config"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

var db *gorm.DB

// Model Model
type Model struct {
	ID int `gorm:"primary_key" json:"id"`
}

// 配置
var conf = *config.Config

// Setup 初始化一个db连接
func Setup() {
	log.Println("数据库初始化")
	db, err := gorm.Open(conf.DbType, fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", conf.DbUser, conf.DbPassword, conf.DbHost, conf.DbPort, conf.DbName))
	if err != nil {
		panic(err)
	}
	db.LogMode(conf.DbLogMode)
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(conf.DbMaxIdleConns)
	db.DB().SetMaxOpenConns(conf.DbMaxOpenConns)
	db.Set("gorm:table_options", "ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;").
		AutoMigrate(&AttendanceDetail{}).AutoMigrate(&Schedule{}).AutoMigrate(&Class{}).
		AutoMigrate(&Detail{}).AutoMigrate(&Group{}).AutoMigrate(&Record{}).
		AutoMigrate(&Schedule{})
}

// GetTx db启动事务
func GetTx() *gorm.DB {
	return db.Begin()
}
