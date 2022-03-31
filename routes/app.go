package routes

import (
	"github.com/fyralabs/id-server/middleware"
	"github.com/fyralabs/id-server/routes/user"
	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App) {
	userGroup := app.Group("/user")
	userGroup.Post("/register", user.Register)
	userGroup.Post("/login", user.Login)

	meGroup := userGroup.Group("/me")
	meGroup.Use(middleware.Auth)
	meGroup.Use("/", middleware.Auth)
	meGroup.Patch("/", user.UpdateMe)
	meGroup.Get("/", user.GetMe)
	meGroup.Get("/sessions", user.GetSessions)
}
