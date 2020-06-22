package main

import (
	"github.com/combataran/taskAPI/config"
	"github.com/combataran/taskAPI/migration"
	"github.com/combataran/taskAPI/route"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func init() {
	db := config.Init()
	migration.Migrate(db)
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := route.SetupRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}
