package model

type Soldier struct {
	ID					uint			`gorm:"primary_key;auto_increment" json:"soldier_id"`
	Name 				string			`gorm:"size:255;not null" json:"soldier_name"`
	MagazinePacks		MagazinePack	`json:"magazine_packs"`
	MagazinePacksValue	string			`json:"-"`
}