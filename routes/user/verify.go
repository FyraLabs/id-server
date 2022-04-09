package user

import (
	"encoding/json"
	"fmt"

	"github.com/fyralabs/id-server/config"
	"github.com/fyralabs/id-server/database"
	"github.com/fyralabs/id-server/ent"
	"github.com/fyralabs/id-server/util"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

// TODO: Rate limit
func RequestVerificationEmail(c *fiber.Ctx) error {
	user := c.Locals("user").(*ent.User)

	if err := util.SendVerificationEmail(user); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

type VerifyEmailBody struct {
	Token string `json:"token" validate:"required"`
}

func VerifyEmail(c *fiber.Ctx) error {
	var verifyData VerifyEmailBody
	if err := json.Unmarshal(c.Request().Body(), &verifyData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid JSON"})
	}

	if err := validate.Struct(verifyData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	token, err := jwt.Parse(verifyData.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
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

	userString, ok := claims["sub"].(string)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"message": "Invalid token"})
	}

	tokenType, ok := claims["type"].(string)
	if !ok || tokenType != "emailVerification" {
		return c.Status(401).JSON(fiber.Map{"message": "Invalid token"})
	}

	emailString, ok := claims["email"].(string)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"message": "Invalid token"})
	}

	parse, err := uuid.Parse(userString)
	if err != nil {
		return err
	}

	u, err := database.DatabaseClient.User.Get(c.Context(), parse)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"message": "Invalid token"})
	}

	if u.Email != emailString {
		return c.Status(401).JSON(fiber.Map{"message": "Expired Token"})
	}

	if _, err := u.Update().SetEmailVerified(true).Save(c.Context()); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}
