package main

import (
	//"fmt"
	//"io"
	"net/http"
	"os"

	//"strconv"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)

	router := gin.New()

	// Add the logger middleware
	router.Use(logger.SetLogger())

	router.GET("/ping", func(c *gin.Context) {
		log.Info("Received ping message")
		c.JSON(http.StatusOK, gin.H{
			"man": "pong",
		})
	})

	//routes.InitRoutes(router)

	log.Info("Port", "9595", "Starting web server")
	router.Run(":9595")
}
