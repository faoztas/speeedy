package db

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"

	"github.com/speeedy/config"
)

var DB *gorm.DB

func Init() {
	db, err := gorm.Open("sqlite3", config.Env.DB.DBFile)
	if err != nil {
		log.Fatal(err)
	}

	if config.Env.DB.Debug {
		db.LogMode(true)
	}

	DB = db
}

func Close() {
	err := DB.Close()
	if err != nil {
		log.Fatal(err)
	}
}
