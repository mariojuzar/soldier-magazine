package service

import (
	"github.com/mariojuzar/soldier-magazine/api/entity/model"
	"github.com/mariojuzar/soldier-magazine/api/libraries/exception"
	"github.com/mariojuzar/soldier-magazine/api/libraries/util"
)

type SoldierService interface {
	CreateSoldier(name string) (model.Soldier, error)
	GetSoldier(id uint) (model.Soldier, error)
	UpdateSoldier(id uint, name string) (model.Soldier, error)
	GetAllSoldier() ([]model.Soldier, error)
	DeleteSoldier(id uint) (model.Soldier, error)
}

type soldierService struct {

}

func NewSoldierService() SoldierService {
	return soldierService{}
}

func (s soldierService) CreateSoldier(name string) (model.Soldier, error) {
	tx := database.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.RollbackUnlessCommitted()
		}
	}()

	gun := model.Gun{
		Name:      name,
		Magazines: []model.Magazine{},
	}

	gunValue := util.ObjectToString(gun)

	soldier := model.Soldier{
		Name:  name,
		Gun:  gun,
		GunValue: gunValue,
	}

	tx.Create(&soldier)

	if err := tx.GetErrors(); len(err) > 0 {
		tx.Rollback()
		return model.Soldier{}, err[0]
	}

	tx.Commit()
	return soldier, nil
}

func (s soldierService) GetSoldier(id uint) (model.Soldier, error) {
	var soldier model.Soldier
	database.DB.Model(model.Soldier{}).Find(&soldier, "id = ?", id)

	if id == soldier.ID {
		if err := database.DB.GetErrors(); len(err) > 0 {
			return model.Soldier{}, err[0]
		}

		return soldier, nil
	} else {
		return model.Soldier{}, exception.NotExistException()
	}
}

func (s soldierService) GetAllSoldier() ([]model.Soldier, error) {
	var soldiers []model.Soldier
	database.DB.Find(&soldiers)

	if err := database.DB.GetErrors(); len(err) > 0 {
		return []model.Soldier{}, err[0]
	}

	return soldiers, nil
}

func (s soldierService) UpdateSoldier(id uint, name string) (model.Soldier, error) {
	soldier, err := s.GetSoldier(id)
	if err != nil {
		return model.Soldier{}, err
	}
	
	soldier.Name = name
	database.DB.Model(model.Soldier{}).Where("id = ?", id).Updates(soldier)

	if err := database.DB.GetErrors(); len(err) > 0 {
		return model.Soldier{}, err[0]
	}
	return soldier, nil
}

func (s soldierService) DeleteSoldier(id uint) (model.Soldier, error) {
	soldier, err := s.GetSoldier(id)
	if err != nil {
		return model.Soldier{}, err
	}

	database.DB.Delete(&soldier)

	if err := database.DB.GetErrors(); len(err) > 0 {
		return model.Soldier{}, err[0]
	}
	return soldier, nil
}