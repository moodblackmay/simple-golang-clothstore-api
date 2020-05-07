package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//Configure DB
func SetupDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "cloth.db")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Cloth{})

	return db
}
