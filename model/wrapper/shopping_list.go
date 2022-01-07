package wrapper

import (
	"shoppinglist/model"

	"gorm.io/gorm"
)

type ShoppingListWrapper struct {
	DB *gorm.DB
}

func CreateShoppingListWrapper(db *gorm.DB) *ShoppingListWrapper {
	return &ShoppingListWrapper{
		DB: db,
	}
}

func (s *ShoppingListWrapper) Insert(data interface{}) error {
	shoppingList := data.(*model.ShoppingList)

	return s.DB.Create(&shoppingList).Error
}

func (s *ShoppingListWrapper) Update(data interface{}) error {
	shoppingList := data.(*model.ShoppingList)

	return s.DB.Save(&shoppingList).Error
}

func (s *ShoppingListWrapper) Delete(data interface{}) error {
	shoppingList := data.(*model.ShoppingList)

	return s.DB.Delete(&shoppingList).Error
}
