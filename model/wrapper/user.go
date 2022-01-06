package wrapper

import (
	"errors"
	"shoppinglist/model"

	"gorm.io/gorm"
)

type UserWrapper struct {
	DB *gorm.DB
}

func CreateUserWrapper(db *gorm.DB) *UserWrapper {
	return &UserWrapper{
		DB: db,
	}
}

func (u *UserWrapper) Insert(data interface{}) error {
	user := data.(*model.User)

	return u.DB.Create(&user).Error
}

func (u *UserWrapper) Update(data interface{}) error {
	user := data.(*model.User)

	return u.DB.Save(&user).Error
}

func (u *UserWrapper) GetUser(userID string) (model.User, error) {
	var user model.User

	err := u.DB.Where(&model.User{UserID: userID}).First(&user).Error
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *UserWrapper) GetAllUser() []string {
	var users []model.User

	userIDs := make([]string, 0)

	err := u.DB.Find(&users).Error
	if err != nil {
		return userIDs
	}

	for _, user := range users {
		userIDs = append(userIDs, user.UserID)
	}

	return userIDs
}

func (u *UserWrapper) UserAlreadyExists(userID string) bool {
	var user model.User
	err := u.DB.Where(&model.User{UserID: userID}).First(&user).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

func (u *UserWrapper) GetShoppingLists(userID string) []model.ShoppingList {
	user, err := u.GetUser(userID)
	if err != nil {
		return nil
	}

	u.DB.Preload("ShoppingLists").Find(&user)
	for i := 0; i < len(user.ShoppingList); i++ {
		u.DB.Preload("Users").Find(&user.ShoppingList[i])
		u.DB.Preload("Items").Find(&user.ShoppingList[i])
	}

	return user.ShoppingList
}
