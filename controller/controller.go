package controller

import (
	"shoppinglist/model/wrapper"

	"gorm.io/gorm"
)

var (
	db          *gorm.DB
	userWrapper *wrapper.UserWrapper
)

// InitializeController initializes the controller with a DB as parameter
func InitializeController(DB *gorm.DB) {
	db = DB
	userWrapper = wrapper.CreateUserWrapper(db)
}
