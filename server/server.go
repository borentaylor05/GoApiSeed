package server

import (
	"log"

	"github.com/borentaylor05/streamline/server/routes"
	"github.com/borentaylor05/streamline/util"
	"github.com/gin-gonic/gin"
)

func HelloHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"text": "Hello World.",
	})
}

func Init() {
	config, err := util.ReadConfig()
	if err != nil {
		log.Printf("ERROR: Unable to parse config file.")
	}
	gin.SetMode(gin.DebugMode)
	if config.Server.GinMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := routes.CreateRouter()

	log.Printf("Starting server on port " + config.Server.Port)
	log.Fatal(router.Run(":" + config.Server.Port))
}
