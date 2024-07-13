package models

type Student struct {
	NPM   string `json:"npm" gorm:"primaryKey"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	City  string `json:"city"`
	Major string `json:"major"`
}
