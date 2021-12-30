package controller

import (
	"net/http"
	"shoppinglist/model"
	//"shoppinglist/model"
	//"shoppinglist/util"

	//"shoppinglist/model/wrapper"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type ShopController struct{}

func (s *ShopController) AddShop(context *gin.Context) {
	//take the input
	inputShop := struct {
		
		Title string `json:"title"`
		
	}{}

	if err := context.BindJSON(&inputShop); err != nil {
		log.Error("Error parsing json input while registering the user", err.Error())
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Error is request body",
		})
		return
	}
	// //if the user already exists
	 var newShop model.ShoppingList
	 newShop.Title = inputShop.Title
	// newUser.Email = inputUser.Email
	// newUser.Name = inputShop.Name
	// encryptedPassword, err := util.HashPassword(inputUser.Password)
	// if err != nil {
	// 	log.Error("Error creating user for user id:", newUser.UserID, err.Error())
	// 	context.JSON(http.StatusInternalServerError, gin.H{
	// 		"status":  http.StatusInternalServerError,
	// 		"message": "Internal Server Error",
	// 	})
	// 	return
	// }
	// newUser.Password = encryptedPassword

	// if userWrapper.UserAlreadyExists(newUser.UserID) {
	// 	log.Error("Error creating user for user id:", newUser.UserID, "User already exists")
	// 	context.JSON(http.StatusConflict, gin.H{
	// 		"status":  http.StatusConflict,
	// 		"message": "User already exists",
	// 	})
	// 	return
	// }
	//create account if new user
	err := shopWrapper.Insert(&newShop)
	if err != nil {
		log.Error("Error creating user for user id:", newShop.Title, err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "User already exists",
		})
		return
	}

	// if newUser.ID == 0 {
	// 	log.Error("Error creating user for user id:", newUser.UserID)
	// 	context.JSON(http.StatusInternalServerError, gin.H{
	// 		"status":  http.StatusInternalServerError,
	// 		"message": "Internal Server Error",
	// 	})
	// 	return
	// }

	//all checks passed
	//log.Infof("User successfully created user ID:%d, for userID:%s", newShop.ID, newShop.UserID)
	context.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": " Shopping List created successfully",
	})
}


