package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID        string         `gorm:"type:varchar(30);not null"`
	Name          string         `gorm:"not null"`
	Email         string         `gorm:"not null"`
	Password      string         `gorm:"type:varchar(80);not null"`
	ShoppingLists []ShoppingList `gorm:"many2many:user_shopping_lists;foreignkey:UserID"`
}
