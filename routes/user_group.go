package routes

import "github.com/gin-gonic/gin"

func UserGroup(r *gin.Engine) {
	userGroup := r.Group("/user")
	{
		userGroup.POST("signup", signupController.AddUser)
	}
}

