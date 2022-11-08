package models

import "gorm.io/gorm"

//create schema for user
type User struct {
	gorm.Model
	Id       int    `json:"id" gorm:"primary_key"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Gender   string `json:"gender"`
}
