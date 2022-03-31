package user

import (
	"github.com/fyralabs/id-server/ent"
	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

func GetSessions(c *fiber.Ctx) error {
	user := c.Locals("user").(*ent.User)

	sessions, err := user.QuerySessions().All(c.Context())
	if err != nil {
		return err
	}

	return c.JSON(lo.Map(sessions, func(s *ent.Session, _ int) fiber.Map {
		return fiber.Map{
			"id":         s.ID.String(),
			"ip":         s.IP,
			"userAgent":  s.UserAgent,
			"createdAt":  s.CreatedAt.String(),
			"lastUsedAt": s.LastUsedAt.String(),
		}
	}))
}
