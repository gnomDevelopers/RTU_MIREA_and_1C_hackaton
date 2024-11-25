package main

import (
	"fmt"
	"log"
	"server/internal/config"
	"server/internal/handler"
	logger "server/internal/log"
	"server/internal/repository/postgres"
	"server/internal/service"

	"github.com/joho/godotenv"
)

// @title VUZ+ API
// @version 1.0
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error while loading .env file: " + err.Error() + "\nMaybe wrong path?")
	}
	conf := config.NewConfig()
	log := logger.InitLogger(conf)
	dbConn, err := postgres.NewDatabase(conf)
	if err != nil {
		log.Fatal().Msg(fmt.Sprintf("could not initialize database connection: %s", err))
	}

	repository := postgres.NewRepository(dbConn.GetDB())
	services := service.NewService(repository, conf)
	//err = services.UserData.AddAdmin(context.Background())
	//if err != nil {
	//	log.Fatal().Msg(err.Error())
	//}
	handlers := handler.NewHandler(services, log, conf)

	app := handlers.Router()
	app.Listen(":8080")
}
