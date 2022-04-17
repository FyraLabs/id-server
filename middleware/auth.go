package middleware

import (
	"time"

	"github.com/fyralabs/id-server/database"
	"github.com/fyralabs/id-server/util"
	"github.com/google/uuid"

	"github.com/gofiber/fiber/v2"
)

func Auth(c *fiber.Ctx) error {
	tokenString, ok := c.GetReqHeaders()["Authorization"]

	if !ok {
		return c.Status(401).JSON(fiber.Map{"message": "Unauthorized"})
	}

	claims, err := util.DecodeJWT(tokenString, "session")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid token"})
	}

	sessionIDString, ok := claims["sub"].(string)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"message": "Invalid token"})
	}

	parse, err := uuid.Parse(sessionIDString)
	if err != nil {
		return err
	}

	s, err := database.DatabaseClient.Session.Get(c.Context(), parse)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"message": "Invalid token"})
	}

	userAgent, ok := c.GetReqHeaders()["User-Agent"]
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "User-Agent header not found"})
	}

	s, err = s.Update().
		SetIP(c.IP()).
		SetUserAgent(userAgent).
		SetLastUsedAt(time.Now()).
		Save(c.Context())

	if err != nil {
		return err
	}

	user, err := s.QueryUser().Only(c.Context())
	if err != nil {
		return err
	}

	c.Locals("session", s)
	c.Locals("user", user)

	return c.Next()
}
