package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"server/internal/config"
	"server/internal/handler"
	logger "server/internal/log"
	"server/internal/repository/postgres"
	"server/internal/service"
)

// @title VUZ+ API
// @version 1.0
// @BasePath /
func main() {
	if err := godotenv.Load("backend/.env"); err != nil {
		log.Fatal("No .env file found")
	}
	conf := config.NewConfig()
	log := logger.InitLogger(conf)
	dbConn, err := postgres.NewDatabase(conf)
	if err != nil {
		log.Fatal().Msg(fmt.Sprintf("could not initialize database connection: %s", err))
	}

	repository := postgres.NewRepository(dbConn.GetDB())
	services := service.NewService(repository, conf)
	handlers := handler.NewHandler(services, log, conf)

	app := handlers.Router()
	app.Listen(":8080")
}
