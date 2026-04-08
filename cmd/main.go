package main

import (
	"gojob/config"
	"gojob/internal/home"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	config.Init() // Starting from .env
	dbConf := config.NewDatabaseConfig()
	log.Println(dbConf)
	app := fiber.New()

	home.NewHandler(app)

	app.Listen(":3000")
}
