package main

import (
	//"fmt"
	//"io"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/viper"
	//"strconv"
	"shoppinglist/config"

	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

func main() {
	//Loading configuration
	config.LoadConfig()
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.Level(viper.GetInt("logging.level")))
	log.Debug("this is a test log")

	router := gin.New()

	// Add the logger middleware
	router.Use(logger.SetLogger())

	router.GET("/ping", func(c *gin.Context) {
		log.Info("Received ping message")
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	//routes.InitRoutes(router)
	port := viper.GetInt("webserver.port")
	log.Info("Port", port, "Starting web server")
	router.Run(fmt.Sprintf(":%d", port))
}
