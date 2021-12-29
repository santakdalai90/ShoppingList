package controller

import (
	"net/http"
	"shoppinglist/model"
	"shoppinglist/util"

	//"shoppinglist/model/wrapper"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type SignupController struct{}

func (s *SignupController) AddUser(context *gin.Context) {
	//take the input
	inputUser := struct {
		UserID   string `json:"user_id"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := context.BindJSON(&inputUser); err != nil {
		log.Error("Error parsing json input while registering the user", err.Error())
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Error is request body",
		})
		return
	}
	//if the user already exists
	var newUser model.User
	newUser.UserID = inputUser.UserID
	newUser.Email = inputUser.Email
	newUser.Name = inputUser.Name
	encryptedPassword, err := util.HashPassword(inputUser.Password)
	if err != nil {
		log.Error("Error creating user for user id:", newUser.UserID, err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Internal Server Error",
		})
		return
	}
	newUser.Password = encryptedPassword

	if userWrapper.UserAlreadyExists(newUser.UserID) {
		log.Error("Error creating user for user id:", newUser.UserID, "User already exists")
		context.JSON(http.StatusConflict, gin.H{
			"status":  http.StatusConflict,
			"message": "User already exists",
		})
		return
	}
	//create account if new user
	err = userWrapper.Insert(&newUser)
	if err != nil {
		log.Error("Error creating user for user id:", newUser.UserID, err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "User already exists",
		})
		return
	}

	if newUser.ID == 0 {
		log.Error("Error creating user for user id:", newUser.UserID)
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Internal Server Error",
		})
		return
	}

	//all checks passed
	log.Infof("User successfully created user ID:%d, for userID:%s", newUser.ID, newUser.UserID)
	context.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "User successfully created",
	})
}
