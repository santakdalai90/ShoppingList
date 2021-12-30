package controller

import (
	"shoppinglist/model/wrapper"

	"gorm.io/gorm"
)

var (
	db          *gorm.DB
	userWrapper *wrapper.UserWrapper
	shopWrapper *wrapper.ShopWrapper

)

func InitializeController(DB *gorm.DB) {
	db = DB
	userWrapper = wrapper.CreateUserWrapper(db)
	shopWrapper = wrapper.CreateShopWrapper(db)
}
