package global

import "gorm.io/gorm"

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func InitDB(d *gorm.DB) {
	db = d
}

