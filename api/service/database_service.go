package service

import (
	"github.com/jinzhu/gorm"
	"github.com/mariojuzar/soldier-magazine/api/entity/model"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type DatabaseService struct {
	DB            *gorm.DB
	IsInitialized bool
}

var databaseService DatabaseService

func Initialize() {
	db, err := gorm.Open("sqlite3", "./api/db/soldier_magazine.db")
	if err != nil {
		log.Fatal("Failed to init db: ", err)
	}
	db.LogMode(true)

	databaseService = DatabaseService{DB: db, IsInitialized:true}

	db.AutoMigrate(
		&model.Magazine{},
		&model.Gun{},
		&model.Soldier{})
}