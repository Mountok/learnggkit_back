package models


type Profile struct {
	Id int `json:"id"`
	UserId int `json:"user_id"`
	Description string `json:"description"`
	Phone string `json:"phone"`
	FullName string `json:"full_name"`
	Image string `json:"image"`
	Score int `json:"score"`	
}

