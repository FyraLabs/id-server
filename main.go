package main

import (
	"github.com/fyralabs/id-server/database"
	"github.com/fyralabs/id-server/routes"
	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database.InitializeDatabase()
	defer database.DatabaseClient.Close()
	app := fiber.New()
	routes.Register(app)
	app.Listen(":3000")
}
