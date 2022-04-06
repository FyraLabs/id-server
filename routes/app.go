package routes

import (
	"github.com/fyralabs/id-server/middleware"
	"github.com/fyralabs/id-server/routes/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Register(app *fiber.App) {
	app.Use(cors.New())

	userGroup := app.Group("/user")
	userGroup.Post("/register", user.Register)
	userGroup.Post("/login", user.Login)
	userGroup.Post("/verifyEmail", user.VerifyEmail)

	meGroup := userGroup.Group("/me")
	meGroup.Use(middleware.Auth)
	meGroup.Use("/", middleware.Auth)
	meGroup.Patch("/", user.UpdateMe)
	meGroup.Get("/", user.GetMe)
	meGroup.Post("/password", user.UpdateMyPassword)
	meGroup.Get("/session", user.GetSessions)
	meGroup.Delete("/session/:id", user.RevokeSession)
}
