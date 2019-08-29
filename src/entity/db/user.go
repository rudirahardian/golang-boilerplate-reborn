package db

import "time"

type User struct {
	ID 				uint       	`gorm:"primary_key" json:"id"`
	Name 			string     	`gorm:"column:name" json:"name"`
	IDCardNumber 	string     	`gorm:"column:id_card_number" json:"id_card_number"`
	Address 		string     	`gorm:"column:address" json:"address"`
	CreatedAt    	*time.Time 	`gorm:"column:created_at" json:"created_at"`
	UpdatedAt    	*time.Time 	`gorm:"column:updated_at" json:"updated_at"`
	DeletedAt    	*time.Time 	`gorm:"column:deleted_at" json:"deleted_at"`
	
}

func (User) TableName() string {
	return "users"
}
