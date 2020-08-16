package request

type SoldierCreateRequest struct {
	Name	string	`json:"name" binding:"required"`
}

type SoldierUpdateRequest struct {
	ID		uint	`json:"id" binding:"required"`
	Name	string	`json:"name" binding:"required"`
}
