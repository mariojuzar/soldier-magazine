package model

type Soldier struct {
	ID			uint	`gorm:"primary_key;auto_increment" json:"soldier_id"`
	Name 		string	`gorm:"size:255;not null" json:"soldier_name"`
	Gun			Gun		`json:"-"`
	GunValue	string	`json:"gun"`
}