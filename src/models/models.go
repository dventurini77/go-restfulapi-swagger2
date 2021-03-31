package models

// User model
type User struct {
	ID    int    `sql:"primary_key;AUTO_INCREMENT" gorm:"primary_key" json:"id" swaggerignore:"true"`
	Name  string `json:"userName" example:"Davide Venturini"`
	Email string `json:"userEMail" example:"d.venturini@gmail.com"`
}
