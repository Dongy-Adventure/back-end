package main

import (
	"fmt"

	"github.com/Dongy-s-Advanture/back-end/internal/config"
	"github.com/Dongy-s-Advanture/back-end/internal/database"
	routes "github.com/Dongy-s-Advanture/back-end/internal/router"
	"github.com/gin-gonic/gin"
)

func main() {
	conf, err := config.LoadConfig()

	if err != nil {
		panic(fmt.Sprintf("Error loading config: %v", err))
	}

	mongoDB, err := database.InitMongoDatabase(&conf.Db)

	if err != nil {
		panic(fmt.Sprintf("Error connecting mongo: %v", err))
	}

	r := routes.NewRouter(gin.Default(), conf)

	r.Run(mongoDB)
}
