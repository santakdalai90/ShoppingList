package routes

import (
	"shoppinglist/controller"

	"github.com/gin-gonic/gin"
)

var (
	signupController = new(controller.SignupController)
	shopController = new(controller.ShopController)
)

func InitRoutes(r *gin.Engine) {
	UserGroup(r)
	ShopGroup(r)
	

}