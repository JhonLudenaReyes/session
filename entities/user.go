package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserId   int    `json:"userId" gorm:"primaryKey"`
	User     string `json:"user"`
	Password string `json:"password"`
	State    string `json:"state" gorm:"default:'A'"`
}
