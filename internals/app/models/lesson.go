package models

type Lesson struct {
	Id int `json:"id"`
	Upkeep string `json:"upkeep"`
	ThemeId int `json:"theme_id"`
}