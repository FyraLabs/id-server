package routes

import (
	"github.com/fyralabs/id-server/middleware"
	"github.com/fyralabs/id-server/routes/user"
	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App) {
	userGroup := app.Group("/user")
	userGroup.Use("/me", middleware.Auth)
	userGroup.Post("/register", user.Register)
	userGroup.Post("/login", user.Login)
	userGroup.Patch("/me", user.UpdateMe)
	userGroup.Get("/me", user.GetMe)
}
