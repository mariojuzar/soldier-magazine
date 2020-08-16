package model

type Gun struct {
	ID			uint32		`gorm:"primary_key;auto_increment" json:"id"`
	Name 		string		`gorm:"size:255;not null" json:"name"`
	Magazines 	[]Magazine	`json:"magazines"`
}
