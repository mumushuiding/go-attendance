package model

import (
	"fmt"
	"log"
	"strconv"

	"github.com/jinzhu/gorm"
	config "github.com/mumushuiding/go-attendance/config"

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
	var err error // 必须这样写，然会出现db为空指针的情况
	db, err = gorm.Open(conf.DbType, fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", conf.DbUser, conf.DbPassword, conf.DbHost, conf.DbPort, conf.DbName))
	if err != nil {
		panic(err)
	}
	// 启用Logger，显示详细日志
	mode, _ := strconv.ParseBool(conf.DbLogMode)
	db.LogMode(mode)
	db.SingularTable(true)
	idle, err := strconv.Atoi(conf.DbMaxIdleConns)
	if err != nil {
		panic(err)
	}
	db.DB().SetMaxIdleConns(idle)
	open, err := strconv.Atoi(conf.DbMaxOpenConns)
	if err != nil {
		panic(err)
	}
	db.DB().SetMaxOpenConns(open)
	db.Set("gorm:table_options", "ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;").
		AutoMigrate(&Class{})
	// AutoMigrate(&AttendanceDetail{}).AutoMigrate(&Schedule{}).
	// AutoMigrate(&Detail{}).AutoMigrate(&Group{}).AutoMigrate(&Record{})
	db.Model(&Class{}).AddIndex("idx_id", "id")
}

// GetTx db启动事务
func GetTx() *gorm.DB {
	return db.Begin()
}

// GetDB GetDB
func GetDB() *gorm.DB {
	return db
}
