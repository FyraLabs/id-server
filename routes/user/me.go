package user

import (
	"github.com/fyralabs/id-server/ent"
	"github.com/gofiber/fiber/v2"
)

func GetMe(c *fiber.Ctx) error {
	user := c.Locals("user").(*ent.User)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":            user.ID.String(),
		"email":         user.Email,
		"name":          user.Name,
		"emailVerified": user.EmailVerified,
		"avatarURL":     user.AvatarURL,
	})
}
