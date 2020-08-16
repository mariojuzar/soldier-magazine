package model

type Soldier struct {
	ID		uint32	`gorm:"primary_key;auto_increment" json:"id"`
	Name 	string	`gorm:"size:255;not null" json:"name"`
	Gun 	Gun		`json:"gun"`
	GunId	uint32	`gorm:"not null" json:"gun_id"`
}