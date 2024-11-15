package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID        string `json:"ID"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Email     string `json:"Email"`
	Mobile    string `json:"Mobile"`
	City      string `json:"City"`
	Country   string `json:"Country"`
	Scope     string `json:"Scope"`
}

// fmt.Println("Insode User Models")
