package model

import (
	"go-mega-code/config"
	"log"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// SetDB func
func SetDB(database *gorm.DB) {
	db = database
}

// ConnectToDB func
func ConnectToDB() *gorm.DB {
	connectionStr := config.GetMysqlConnectingString()
	log.Println("Connect to db...,")
	db, err := gorm.Open("mysql", connectionStr)
	if err != nil {
		panic("Failed to connect database, error: " + err.Error())
	}
	db.SingularTable(true)
	return db
}
