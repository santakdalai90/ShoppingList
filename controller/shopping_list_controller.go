package controller

import (
	"net/http"
	"shoppinglist/cache"
	"shoppinglist/model"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type ShoppingListController struct{}

func (s *ShoppingListController) AddShoppingList(context *gin.Context) {
	status, resp := cache.AuthorizeSessionToken(context)
	if status != http.StatusOK {
		context.JSON(status, gin.H{
			"status":           status,
			"message":          "User unauthorized",
			"shopping_list_id": 0,
		})
		return
	}
	//take the input
	inputShoppingList := struct {
		Title string `json:"title"`
	}{}

	if err := context.BindJSON(&inputShoppingList); err != nil {
		log.Error("Error parsing json input while creating the shopping list", err.Error())
		context.JSON(http.StatusBadRequest, gin.H{
			"status":           http.StatusBadRequest,
			"message":          "Error is request body",
			"shopping_list_id": 0,
		})
		return
	}

	var newShoppingList model.ShoppingList
	newShoppingList.Title = inputShoppingList.Title

	user, _ := userWrapper.GetUser(string(resp.([]byte)))
	newShoppingList.Users = make([]model.User, 0)
	newShoppingList.Users = append(newShoppingList.Users, user)

	//create account if new user
	err := shoppingListWrapper.Insert(&newShoppingList)
	if err != nil {
		log.Error(
			"Error creating shopping list for user id:",
			user.UserID,
			"Shopping list title:",
			inputShoppingList.Title,
			err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":           http.StatusInternalServerError,
			"message":          "Internal server error",
			"shopping_list_id": 0,
		})
		return
	}

	if newShoppingList.ID == 0 {
		log.Error(
			"Error creating shopping list for user id:",
			user.UserID,
			"Shopping list title:",
			inputShoppingList.Title,
		)
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":           http.StatusInternalServerError,
			"message":          "Internal Server Error",
			"shopping_list_id": 0,
		})
		return
	}

	//all checks passed
	log.Infof("Shopping list successfully created shopping list title:%d, for userID:%s", newShoppingList.Title, user.UserID)
	context.JSON(http.StatusCreated, gin.H{
		"status":           http.StatusCreated,
		"message":          "Shopping list successfully created",
		"shopping_list_id": newShoppingList.ID,
	})
}
