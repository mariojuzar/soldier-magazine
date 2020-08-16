package tests

import (
	"github.com/go-playground/assert/v2"
	"github.com/mariojuzar/soldier-magazine/api/entity/model"
	"log"
	"testing"
)

func TestRefreshTable(t *testing.T) {
	err := refreshAllTable()
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, nil, err)
}

func refreshAllTable() error {
	err := database.DB.DropTableIfExists(&model.Magazine{}, &model.Gun{}, &model.Soldier{}).Error
	if err != nil {
		return err
	}
	err = database.DB.AutoMigrate(&model.Magazine{}, &model.Gun{}, &model.Soldier{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed table")
	return nil
}