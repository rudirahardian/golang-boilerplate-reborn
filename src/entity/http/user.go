package http

import "time"

type UserRequest struct {
	Name 			string     	`json:"name"`
	IDCardNumber 	string     	`json:"id_card_number"`
	Address 		string     	`json:"address"`
}

type UserRequestNew struct {
	Name 			string     	`json:"name"`
	IDCardNumber 	string     	`json:"id_card_number"`
	Address 		string     	`json:"address"`
	UserStatusID	uint		`json:"user_status_id"`
}

type UserDetailResponse struct {
	ID 				uint       	`json:"id"`
	Name 			string     	`json:"name"`
	IDCardNumber 	string     	`json:"id_card_number"`
	Address 		string     	`json:"address"`
	Status 			*string     `json:"status"`
	Avatar 			*string     `json:"avatar"`
	CreatedAt    	*time.Time 	`json:"created_at"`
	UpdatedAt    	*time.Time 	`json:"updated_at"`
}

type UserResponse struct {
	ID 				uint       	`json:"id"`
	Name 			string     	`json:"name"`
	IDCardNumber 	string     	`json:"id_card_number"`
	Address 		string     	`json:"address"`
	CreatedAt    	*time.Time 	`json:"created_at"`
	UpdatedAt    	*time.Time 	`json:"updated_at"`
}