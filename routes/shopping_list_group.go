package routes

import "github.com/gin-gonic/gin"

func ShoppingListGroup(r *gin.Engine) {
	shoppingListGroup := r.Group("/shopping-list")
	{
		shoppingListGroup.POST("create", shoppingListController.AddShoppingList)
	}
}
