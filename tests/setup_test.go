package tests

import (
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/mariojuzar/soldier-magazine/api/entity/model"
	"github.com/mariojuzar/soldier-magazine/api/service"
	"log"
	"os"
	"testing"
)

var database service.DatabaseService

func TestMain(m *testing.M)  {
	var err error
	err = godotenv.Load(os.ExpandEnv(".env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}

	Databases()

	os.Exit(m.Run())
}

func Databases() {
	db, err := gorm.Open("sqlite3", "soldier_magazine.db")
	if err != nil {
		log.Fatal("Failed to init db: ", err)
	}
	db.LogMode(true)

	database = service.DatabaseService{DB: db, IsInitialized:true}

	db.AutoMigrate(
		&model.Magazine{},
		&model.Gun{},
		&model.Soldier{})
}