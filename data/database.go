package data

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB = initDatabase(Config.AutoMigrate)

func initDatabase(automigrateflag bool) *gorm.DB {
	database, err := gorm.Open(sqlite.Open("EncuestifyDb.db"))

	if err != nil {
		panic("Failed to connect to database " + err.Error())
	}

	if automigrateflag {
		database.AutoMigrate()
	}

	return database
}
