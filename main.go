package main

import (
    //"fmt"
	//"io"
	"net/http"
	//"os"
    "github.com/spf13/viper"
    "strconv"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
    zerolog.SetGlobalLevel(zerolog.Level(1))
	// logFile, err := os.OpenFile(viper.GetString("logging.path"), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	// if err != nil {
	// 	log.Fatal().Msgf("Error creating or writing to the log file. %s", err.Error())
	// }
	// defer logFile.Close()

    router := gin.New()

	// Add the logger middleware
	router.Use(logger.SetLogger())

	router.GET("/ping", func(c *gin.Context) {
		log.Info().Msg("Received ping message")
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	//routes.InitRoutes(router)

	log.Info().Str("Port", strconv.Itoa(viper.GetInt("webserver.port"))).Msg("Starting web server")
	//router.Run(fmt.Sprintf(":%s", strconv.Itoa(viper.GetInt("webserver.port"))))
}