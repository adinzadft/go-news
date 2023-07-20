package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DB() *gorm.DB {

	dsn := "zidan1102:2011&@d!$04@tcp(localhost:3306)/go-news?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	println("berhasil")

	return db
}
