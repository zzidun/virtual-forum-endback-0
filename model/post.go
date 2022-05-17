package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Category   *Category
	CategoryId uint   `gorm:"not null"`
	Title      string `gorm:"type:varchar(50);not null"`
	Speak      uint   `gorm:"not null"`
	User       *User
	UserId     uint `gorm:"not null"`
}

type PostPostForm struct {
	CategoryId string `json:"categoryid" binding:"required"`
	Title      string `json:"title" binding:"required"`
	UserId     string `json:"userid" binding:"required"`
}
