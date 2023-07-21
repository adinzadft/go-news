package mysql

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DB() *gorm.DB {

	dsn := os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_HOST") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	println("berhasil")

	return db
}
