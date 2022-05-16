package main

import (
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
		TrustedProxies: []string{"10.42.0.0/16", "103.21.244.0/22", "103.22.200.0/22", "103.31.4.0/22", "104.16.0.0/12", "108.162.192.0/18", "131.0.72.0/22", "141.101.64.0/18", "162.158.0.0/15", "172.64.0.0/13", "173.245.48.0/20", "188.114.96.0/20", "190.93.240.0/20", "197.234.240.0/22", "198.41.128.0/17", "2400:cb00::/32", "2606:4700::/32", "2803:f800::/32", "2405:b500::/32", "2405:8100::/32", "2c0f:f248::/32", "2a06:98c0::/29"},
	})
	routes.Register(app)
	if err := app.Listen(":8080"); err != nil {
		panic(err.Error())
	}
}
