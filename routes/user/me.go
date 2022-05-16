package user

import (
	"github.com/fyralabs/id-server/ent"
	"github.com/gofiber/fiber/v2"
)

func GetMe(c *fiber.Ctx) error {
	println(c.Get("X-Forwarded-For"))
	println(c.IP())
	user := c.Locals("user").(*ent.User)

	if user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "unauthorized",
		})
	}

	var avatarURL *string

	if user.AvatarURL != nil {
		avatarURL = user.AvatarURL
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":            user.ID.String(),
		"email":         user.Email,
		"name":          user.Name,
		"emailVerified": user.EmailVerified,
		"avatarURL":     avatarURL,
	})
}
