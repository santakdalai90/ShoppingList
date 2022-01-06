package controller

import (
	"net/http"
	"shoppinglist/cache"
	"shoppinglist/model"
	"shoppinglist/util"

	//"shoppinglist/model/wrapper"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// SignupController is a controller handling user sign-ins
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

// Login logs in a user to the application based on the credentials passed in the context
func (s *SignupController) Login(context *gin.Context) {
	loginInfo := struct {
		UserID   string `json:"user_id"`
		Password string `json:"password"`
	}{}

	if err := context.BindJSON(&loginInfo); err != nil {
		log.Errorf("Error occurred while reading the input request body. %s", err.Error())
		context.JSON(http.StatusBadRequest, gin.H{
			"status":        http.StatusBadRequest,
			"message":       "Error in the request body",
			"session_token": nil,
		})
		return
	}

	user, err := userWrapper.GetUser(loginInfo.UserID)

	if err != nil {
		log.Errorf("Error occurred while searching for the user %v in the database. %s", loginInfo.UserID, err.Error())
		context.JSON(http.StatusNotFound, gin.H{
			"status":        http.StatusNotFound,
			"message":       "Mentioned user not found",
			"session_token": nil,
		})
		return
	}

	if !util.CheckPasswordHash(loginInfo.Password, user.Password) {
		log.Errorf("Invalid username/email or password for user: %s", loginInfo.UserID)
		context.JSON(http.StatusUnauthorized, gin.H{
			"status":        http.StatusUnauthorized,
			"message":       "Invalid username/email or password",
			"session_token": nil,
		})
		return
	}

	// Create a random session token
	sessionToken, err := cache.CreateSessionToken(user)
	if err != nil {
		log.Error(err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":        http.StatusInternalServerError,
			"message":       "Internal Server Error",
			"session_token": nil,
		})
		return
	}

	log.Infof("User: %s, successfully logged in", loginInfo.UserID)
	context.JSON(http.StatusOK, gin.H{
		"status":        http.StatusOK,
		"message":       "User successfully logged in",
		"session_token": sessionToken,
	})
}

// Logout performs log out based on the session cookie passed in the context
func (s *SignupController) Logout(context *gin.Context) {
	status, resp := cache.AuthorizeSessionToken(context)
	if status != http.StatusOK {
		context.JSON(status, gin.H{
			"status":  status,
			"message": "User unauthorised",
		})
		return
	}

	if !cache.DeleteSessionToken(context) {
		context.JSON(status, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Internal error while logging out",
		})
		return
	}

	log.Infof("User successfully logged out userID: %s", string(resp.([]byte)))
	context.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "User successfully logged out",
	})
}
