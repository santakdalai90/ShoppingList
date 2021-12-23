package model

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Name           string `gorm:"not null"`
	Description    string
	IsBought       bool
	BoughtBy       string `gorm:"foreignkey:BoughtBy;association_foreignkey:UserID"`
	CategoryID     uint
	ShoppingListID uint
}
