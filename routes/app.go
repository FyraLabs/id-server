package routes

import (
	"github.com/fyralabs/id-server/middleware"
	"github.com/fyralabs/id-server/routes/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Register(app *fiber.App) {
	app.Use(cors.New())

	app.Get("/lookup", user.LookupConnectToken)

	userGroup := app.Group("/user")
	userGroup.Post("/register", user.Register)
	userGroup.Post("/login", user.Login)
	userGroup.Post("/login/2fa", user.Login2FA)
	userGroup.Post("/verifyEmail", user.VerifyEmail)

	meGroup := userGroup.Group("/me")
	meGroup.Use(middleware.Auth)
	meGroup.Use("/", middleware.Auth)
	meGroup.Patch("/", user.UpdateMe)
	meGroup.Get("/", user.GetMe)
	meGroup.Put("/avatar", user.UploadAvatar)
	meGroup.Delete("/avatar", user.DeleteAvatar)
	meGroup.Post("/requestVerificationEmail", user.RequestVerificationEmail)
	meGroup.Post("/password", user.UpdateMyPassword)
	meGroup.Get("/session", user.GetSessions)
	meGroup.Delete("/session/:id", user.RevokeSession)
	meGroup.Get("/2fa", user.GetMethods)
	meGroup.Post("/2fa", user.AddMethod)
	meGroup.Delete("/2fa/:id", user.RemoveMethod)
	meGroup.Post("/connect", user.Connect)
}
