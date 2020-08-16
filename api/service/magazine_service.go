package service

import (
	"github.com/mariojuzar/soldier-magazine/api/entity/model"
	"github.com/mariojuzar/soldier-magazine/api/entity/rest-web/request"
	"github.com/mariojuzar/soldier-magazine/api/entity/rest-web/response"
	"github.com/mariojuzar/soldier-magazine/api/libraries/exception"
	"github.com/mariojuzar/soldier-magazine/api/libraries/util"
)

type MagazineService interface {
	PutMagazine(request  request.PutMagazineRequest) (model.Soldier, error)
	LoadMagazineRandomly(request request.LoadMagazineRequest) (response.VerifiedMagazineStatusResponse, error)
}

type magazineService struct {

}

func NewMagazineService() MagazineService {
	return magazineService{}
}

func (m magazineService) PutMagazine(request request.PutMagazineRequest) (model.Soldier, error) {
	soldier, err := getSoldier(request.SoldierID)

	if err != nil {
		return model.Soldier{}, err
	}

	var newMagazine []model.Magazine
	for _, val := range request.Magazines {
		newMagazine = append(newMagazine, model.Magazine{
			Capacity:    0,
			MaxCapacity: val.MaxCapacity,
			IsVerified:  false,
		})
	}

	soldier.MagazinePacks.Magazines = append(soldier.MagazinePacks.Magazines, newMagazine...)

	soldier.MagazinePacksValue = util.ObjectToString(soldier.MagazinePacks)
	database.DB.Model(&model.Soldier{}).Where("id = ?", soldier.ID).Updates(soldier)

	if err := database.DB.GetErrors(); len(err) > 0 {
		return model.Soldier{}, err[0]
	}
	return soldier, nil
}

func getSoldier(id uint) (model.Soldier, error) {
	var soldier model.Soldier
	database.DB.Model(model.Soldier{}).Find(&soldier, "id = ?", id)

	if id == soldier.ID {
		if err := database.DB.GetErrors(); len(err) > 0 {
			return model.Soldier{}, err[0]
		}

		var magazinePack model.MagazinePack
		util.StringToObject(soldier.MagazinePacksValue, &magazinePack)
		soldier.MagazinePacks = magazinePack
		return soldier, nil
	} else {
		return model.Soldier{}, exception.NotExistException()
	}
}

func (m magazineService) LoadMagazineRandomly(request request.LoadMagazineRequest) (response.VerifiedMagazineStatusResponse, error) {
	soldier, err := getSoldier(request.SoldierID)

	if err != nil {
		return response.VerifiedMagazineStatusResponse{}, err
	}
	
	for idx, val := range request.Bullets {
		safeIdx := ensureIdxMagazine(soldier, idx)
		soldier.MagazinePacks.Magazines[safeIdx].Capacity = ensureCapacity(val, soldier.MagazinePacks.Magazines[safeIdx].MaxCapacity)

		if testFiringToGun(&soldier.MagazinePacks.Magazines[safeIdx]) {
			soldier.MagazinePacksValue = util.ObjectToString(soldier.MagazinePacks)
			database.DB.Model(&model.Soldier{}).Where("id = ?", soldier.ID).Updates(&soldier)
			return response.VerifiedMagazineStatusResponse{SoldierID: soldier.ID, IsVerifiedMagazine: true}, nil
		}
	}

	database.DB.Model(&model.Soldier{}).Where("id = ?", soldier.ID).Updates(&soldier)

	return response.VerifiedMagazineStatusResponse{SoldierID: soldier.ID, IsVerifiedMagazine: false}, nil
}

func ensureIdxMagazine(soldier model.Soldier, idx int) int {
	if idx >= len(soldier.MagazinePacks.Magazines) {
		idx = 0
	}
	return idx
}

func ensureCapacity(capacity uint, max uint) uint {
	if capacity >= max {
		return max
	}
	return capacity
}

func testFiringToGun(magazine *model.Magazine) bool {
	if magazine.Capacity == magazine.MaxCapacity {
		magazine.IsVerified = true
	}
	return magazine.IsVerified
}