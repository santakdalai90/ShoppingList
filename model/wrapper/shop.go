package wrapper

import (
	"errors"
	"shoppinglist/model"

	"gorm.io/gorm"
)

type ShopWrapper struct {
	DB *gorm.DB
}

func CreateShopWrapper(db *gorm.DB) *ShopWrapper {
	return &ShopWrapper{
		DB: db,
	}
}

func (s *ShopWrapper) Insert(data interface{}) error {
	shop := data.(*model.ShoppingList)

	return s.DB.Create(&shop).Error
}

func (s *ShopWrapper) Update(data interface{}) error {
	user := data.(*model.ShoppingList)

	return s.DB.Save(&user).Error
}

func (s *ShopWrapper) GetUser(userID string) (model.ShoppingList, error) {
	var user model.ShoppingList

	err := s.DB.Where(&model.ShoppingList{Title: userID}).First(&user).Error
	if err != nil {
		return model.ShoppingList{}, err
	}

	return user, nil
}

func (s *ShopWrapper) GetAllUser() []string {
	var users []model.ShoppingList

	userIDs := make([]string, 0)

	err := s.DB.Find(&users).Error
	if err != nil {
		return userIDs
	}

	for _, user := range users {
		userIDs = append(userIDs, user.Title)
	}

	return userIDs
}

func (s *ShopWrapper) UserAlreadyExists(userID string) bool {
	var user model.User
	err := s.DB.Where(&model.ShoppingList{Title: userID}).First(&user).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

// func (s *ShopWrapper) GetShoppingLists(userID string) []model.ShoppingList {
// 	user, err := s.GetUser(userID)
// 	if err != nil {
// 		return nil
// 	}

// 	s.DB.Preload("ShoppingLists").Find(&user)
// 	for i := 0; i < len(ShoppingList); i++ {
// 		s.DB.Preload("Users").Find(&user.ShoppingList[i])
// 		s.DB.Preload("Items").Find(&user.ShoppingList[i])
// 	}

// 	return user.ShoppingList
// }
