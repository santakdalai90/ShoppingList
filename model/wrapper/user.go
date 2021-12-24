package wrapper

import (
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
