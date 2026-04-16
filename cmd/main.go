package main

import (
	"gojob/config"
	"gojob/internal/home"
	"gojob/pkg/logger"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	config.Init() // Starting from .env
	config.NewDatabaseConfig()
	logConfig := config.NewLogConfig()
	customLogger := logger.NewLogger(logConfig)

	app := fiber.New()

	//log levels
	//log.SetLevel(log.Level(logConfig.Level))
	//zerolog.SetGlobalLevel(zerolog.InfoLevel)
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: customLogger,
	}))
	app.Use(recover.New()) // to process panics, recovery
	app.Static("/public", "./public")

	home.NewHandler(app, customLogger)

	app.Listen(":3000")
}
