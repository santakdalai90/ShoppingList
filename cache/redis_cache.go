package cache

import (
	"fmt"
	"net/http"
	"shoppinglist/model"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var cache redis.Conn

// InitCache initializes redis cache based on the configuration of the app
func InitCache() {
	conn, err := redis.DialURL(fmt.Sprintf("redis://%s", viper.GetString("kv.ip")))
	if err != nil {
		panic(err)
	}
	log.Info("Connection to Redis established")
	cache = conn
}

// CreateSessionToken creates and returns a session token and stores it in redis cache
func CreateSessionToken(user model.User) (string, error) {
	// Create a random session token
	sessionToken := uuid.NewV4().String()
	// Save the token in the redis cache along with the username.
	_, err := cache.Do("SET", sessionToken, user.UserID)
	if err != nil {
		return "", fmt.Errorf("error occurred while saving the token to redis. %s", err.Error())
	}

	return sessionToken, nil
}

// AuthorizeSessionToken reads the token from the context and validates it
func AuthorizeSessionToken(context *gin.Context) (int, interface{}) {
	sessionToken := context.Request.Header.Get("session_token")
	if sessionToken == "" {
		log.Error("Invalid/Bad request. No session token specified")
		return http.StatusBadRequest, nil
	}
	response, err := cache.Do("GET", sessionToken)
	if err != nil {
		log.Errorf("Error occurred while fetching token from redis. %v", err.Error())
		return http.StatusInternalServerError, nil
	}
	if response == nil {
		log.Error("Unauthorised access attempt or invalid token")
		return http.StatusUnauthorized, nil
	}

	log.Infof("Successful authentication using the token: %v", sessionToken)
	return http.StatusOK, response
}

// DeleteSessionToken deletes the session token as provided in the context
func DeleteSessionToken(context *gin.Context) bool {
	sessionToken := context.Request.Header.Get("session_token")

	response, err := cache.Do("DEL", sessionToken)
	if err != nil {
		log.Errorf("Error occurred while deleting token from redis. %v", err.Error())
		return false
	}
	if response == nil {
		log.Error("Unauthorised access attempt or invalid token")
		return false
	}

	log.Infof("Successful logout using the token: %v", sessionToken)
	return true
}
