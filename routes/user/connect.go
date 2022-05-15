package user

import (
	"time"

	"github.com/fyralabs/id-server/config"
	"github.com/fyralabs/id-server/database"
	"github.com/fyralabs/id-server/ent"
	"github.com/fyralabs/id-server/util"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func Connect(c *fiber.Ctx) error {
	session := c.Locals("session").(*ent.Session)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  session.ID,
		"type": "connect",
	})

	tokenString, err := token.SignedString([]byte(config.Environment.JwtKey))

	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"token": tokenString,
	})
}

type ConnectInput struct {
	Token       string `json:"token" validate:"required,jwt"`
	IP          string `json:"ip" validate:"required"`
	UserAgent   string `json:"userAgent" validate:"required"`
	ClientToken string `json:"clientToken" validate:"required,jwt"`
}

func LookupConnectToken(c *fiber.Ctx) error {
	var requestData ConnectInput

	if err := c.QueryParser(&requestData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid JSON"})
	}

	if err := validate.Struct(requestData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	claims, err := util.DecodeJWT(requestData.Token, "connect")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid token"})
	}

	clientClaims, err := util.DecodeJWT(requestData.ClientToken, "client")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid client token"})
	}

	clientIDString, ok := clientClaims["sub"].(string)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"message": "Invalid client token"})
	}

	_, ok = util.ConnectClientIDs[clientIDString]
	if !ok {
		return c.Status(401).JSON(fiber.Map{"message": "Invalid client token"})
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

	s, err = s.Update().
		SetIP(requestData.IP).
		SetUserAgent(requestData.UserAgent).
		SetLastUsedAt(time.Now()).
		Save(c.Context())

	if err != nil {
		return err
	}

	user, err := s.QueryUser().Only(c.Context())
	if err != nil {
		return err
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
