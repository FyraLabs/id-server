package main

import (
	"github.com/fyralabs/id-server/config"
	"github.com/fyralabs/id-server/database"
	"github.com/fyralabs/id-server/routes"
	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	err := config.InitializeEnv()
	if err != nil {
		panic(err.Error())
	}

	err = database.InitializeDatabase()
	if err != nil {
		panic(err.Error())
	}
	defer database.DatabaseClient.Close()

	app := fiber.New()
	routes.Register(app)
	if err := app.Listen(":3000"); err != nil {
		panic(err.Error())
	}
}