package user

import (
	"github.com/fyralabs/id-server/database"
	"github.com/fyralabs/id-server/ent"
	"github.com/fyralabs/id-server/ent/session"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

func RevokeSession(c *fiber.Ctx) error {
	user := c.Locals("user").(*ent.User)

	sessionIdString := c.Params("id")

	sessionId, err := uuid.Parse(sessionIdString)
	if err != nil {
		return err
	}

	// This should prevent someone from revoking someone else's session
	session, err := user.QuerySessions().Where(session.ID(sessionId)).Only(c.Context())
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "Session not found"})
	}

	if err := database.DatabaseClient.Session.DeleteOne(session).Exec(c.Context()); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}
