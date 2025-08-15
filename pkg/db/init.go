package db

import (
	"Blug/pkg/entities"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var _db *gorm.DB

func SqliteInit() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("Blug.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	autoMigrate(db)
	return db
}

func autoMigrate(db *gorm.DB) {
	err := db.Set(
		"gorm:table_settings", "charset=utf8&parseTime=True&loc=Local").
		AutoMigrate(&entities.Article{}, &entities.User{}, &entities.Class{})
	if err != nil {
		panic(err)
	}
}
