package model

type Magazine struct {
	Capacity	uint 	`json:"capacity"`
	MaxCapacity	uint	`json:"max_capacity"`
	IsVerified	bool	`json:"is_verified"`
}
