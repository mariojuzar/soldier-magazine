package request

type SoldierCreateRequest struct {
	Name	string	`json:"name" binding:"required"`
}

type SoldierUpdateRequest struct {
	ID		uint	`json:"id" binding:"required"`
	Name	string	`json:"name" binding:"required"`
}

type MagazineRequest struct {
	MaxCapacity	uint	`json:"max_capacity" binding:"required"`
}

type PutMagazineRequest struct {
	SoldierID 	uint				`json:"soldier_id" binding:"required"`
	Magazines 	[]MagazineRequest	`json:"magazines" binding:"required"`
}

type LoadMagazineRequest struct {
	SoldierID 	uint	`json:"soldier_id" binding:"required"`
	Bullets		[]uint	`json:"bullets"`
}
