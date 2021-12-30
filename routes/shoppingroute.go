package routes

import "github.com/gin-gonic/gin"

func ShopGroup(r *gin.Engine) {
	shopGroup := r.Group("/user")
	{
		shopGroup.POST("shoppinglist", shopController.AddShop)
	}
}

