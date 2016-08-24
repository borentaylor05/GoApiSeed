package server

import (
	"GoApiSeed/server/routes"
	"GoApiSeed/util"
	"log"
)

func Init() {
	router := routes.CreateRouter()

	config, err := util.ReadConfig()
	if err != nil {
		log.Printf("ERROR: Unable to parse config file.")
	}

	log.Printf("Starting server on port " + config.Server.Port)
	log.Fatal(router.Run(":" + config.Server.Port))
}
