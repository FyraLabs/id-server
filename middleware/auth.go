package middleware

import (
	"fmt"
	"github.com/fyralabs/id-server/database"
	"github.com/google/uuid"
	"time"

	"github.com/fyralabs/id-server/config"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Auth(c *fiber.Ctx) error {
	tokenString, ok := c.GetReqHeaders()["Authorization"]

	if !ok {
		return c.Status(401).JSON(fiber.Map{"message": "Unauthorized"})
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.Environment.JwtKey), nil
	})

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Cannot parse token"})
	}

	if !token.Valid {
		return c.Status(401).JSON(fiber.Map{"message": "Invalid token"})
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return c.Status(400).JSON(fiber.Map{"message": "Cannot parse token"})
	}

	if err := claims.Valid(); err != nil {
		return c.Status(401).JSON(fiber.Map{"message": "Invalid token"})
	}

	sessionString, ok := claims["sub"].(string)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"message": "Invalid token"})
	}

	tokenType, ok := claims["type"].(string)
	if !ok || tokenType != "session" {
		return c.Status(401).JSON(fiber.Map{"message": "Invalid token"})
	}

	parse, err := uuid.Parse(sessionString)
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
