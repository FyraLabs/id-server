package main

import (
	"strings"

	"github.com/fyralabs/id-server/config"
	"github.com/fyralabs/id-server/database"
	"github.com/fyralabs/id-server/routes"
	"github.com/fyralabs/id-server/util"
	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	if err := config.InitializeEnv(); err != nil {
		panic(err.Error())
	}

	util.InitializeConnectClients()

	if err := util.InitializeS3(); err != nil {
		panic(err.Error())
	}

	util.InitializeSendGrid()

	if err := util.InitializeGeoIP(); err != nil {
		panic(err.Error())
	}
	defer util.GeoIP.Close()

	if err := database.InitializeDatabase(); err != nil {
		panic(err.Error())
	}

	defer database.DatabaseClient.Close()

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			c.Status(500).JSON(fiber.Map{"message": "Internal server error"})
			return nil
		},
		EnableTrustedProxyCheck: true,
		TrustedProxies:          strings.Split(config.Environment.TrustedProxies, ","),
	})
	routes.Register(app)
	if err := app.Listen(":8080"); err != nil {
		panic(err.Error())
	}
}
