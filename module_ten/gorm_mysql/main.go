package main

import (
	"github.com/hjtcn/benbird-and-annie-go/pkg/apis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func connectDb() (db *gorm.DB) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_learning?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("数据库连接失败")
	}

	return db
}

func selectDb(db *gorm.DB) {
	res := []*apis.PersonalInfo{}
}

func main() {
	db := connectDb()

}
