package service

import (
	"github.com/jinzhu/gorm"
	"github.com/mariojuzar/soldier-magazine/api/entity/model"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Database struct {
	DB            *gorm.DB
	IsInitialized bool
}

type DatabaseService interface {
	Initialize() error
}

type databaseService struct {

}

func NewDatabaseService() DatabaseService {
	return databaseService{}
}

var database Database

func (d databaseService) Initialize() error {
	db, err := gorm.Open("sqlite3", "./api/db/soldier_magazine.db")
	if err != nil {
		log.Fatal("Failed to init db: ", err)
	}
	db.LogMode(true)

	database = Database{DB: db, IsInitialized:true}

	db.AutoMigrate(&model.Soldier{})

	return nil
}