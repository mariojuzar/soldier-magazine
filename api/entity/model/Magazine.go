package model

type Magazine struct {
	ID			uint32	`gorm:"primary_key;auto_increment" json:"id"`
	Capacity	uint32 	`gorm:"not null" json:"capacity"`
	MaxCapacity	uint32	`gorm:"not null" json:"max_capacity"`
}
