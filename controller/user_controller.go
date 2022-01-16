package controller

import (
	"net/http"
	"shoppinglist/cache"
	"shoppinglist/model"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type UserController struct{}

func (u *UserController) FetchShoppingLists(context *gin.Context) {
	status, resp := cache.AuthorizeSessionToken(context)
	if status != http.StatusOK {
		context.JSON(status, gin.H{
			"status":        status,
			"message":       "User unauthorized",
			"shopping_list": []model.ShoppingList{},
		})
		return
	}

	userID := string(resp.([]byte))
	shoppingLists := userWrapper.GetShoppingLists(userID)

	if len(shoppingLists) == 0 {
		log.Warnf("No shopping list found for user id %s", userID)
		context.JSON(http.StatusNotFound, gin.H{
			"status":        http.StatusNotFound,
			"message":       "Shopping list not found for the user",
			"shopping_list": []model.ShoppingList{},
		})
		return
	}

	log.Info("Shopping lists fetched successfully for the user id %s", userID)
	context.JSON(http.StatusOK, gin.H{
		"status":        http.StatusOK,
		"message":       "Shopping lists found for the user",
		"shopping_list": shoppingLists,
	})
}
