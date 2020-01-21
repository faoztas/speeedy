package db

import (
	"log"

	"github.com/boltdb/bolt"

	"speedy/config"
)

var DB *bolt.DB

func Init() {
	db, err := bolt.Open(config.Env.DBFile, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	DB = db
}

func Close() {
	err := DB.Close()
	if err != nil{
		log.Fatal(err)
	}
}
