package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/spf13/viper"

	"shoppinglist/config"

	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() {
	var err error

	var dbConfig = struct {
		username             string
		password             string
		dbname               string
		ip                   string
		port                 int
		maxConnectionAttempt int
	}{
		username:             viper.GetString("database.username"),
		password:             viper.GetString("database.password"),
		dbname:               viper.GetString("database.dbname"),
		ip:                   viper.GetString("database.ip"),
		port:                 viper.GetInt("database.port"),
		maxConnectionAttempt: viper.GetInt("database.max_connection_attempt"),
	}

	dbConnStr := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.username,
		dbConfig.password,
		dbConfig.ip,
		dbConfig.port,
		dbConfig.dbname,
	)

	var attemptCount int
	for attemptCount = 0; attemptCount < dbConfig.maxConnectionAttempt; attemptCount++ {
		db, err = gorm.Open(mysql.Open(dbConnStr), &gorm.Config{})
		if err == nil {
			break
		}
		fmt.Printf("Attempting connection to database. Attemp count: %d\n", attemptCount)
		time.Sleep(5 * time.Second)
	}

	if attemptCount == dbConfig.maxConnectionAttempt {
		panic("failed to connect to database. Error" + err.Error())
	}

	fmt.Println("Database connected successfully")
}

func main() {
	//Loading configuration
	config.LoadConfig()

	//initializing database
	initDB()

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
