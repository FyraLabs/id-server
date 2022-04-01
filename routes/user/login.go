package user

import (
	"encoding/json"
	"github.com/fyralabs/id-server/config"

	"github.com/fyralabs/id-server/database"
	"github.com/fyralabs/id-server/ent/user"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/matthewhartstonge/argon2"
)

type LoginUser struct {
	Email    string `json:"email" validate:"required,email,min=5,max=256"`
	Password string `json:"password" validate:"required,min=8,max=256"`
}

func Login(c *fiber.Ctx) error {
	body := c.Request().Body()

	var userData LoginUser
	if err := json.Unmarshal(body, &userData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid JSON"})
	}

	if err := validate.Struct(userData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	exists, err := database.DatabaseClient.User.Query().Where(user.EmailEQ(userData.Email)).Exist(c.Context())
	if err != nil {
		return err
	}

	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User doesn't exist"})
	}

	u, err := database.DatabaseClient.User.
		Query().
		Where(user.EmailEQ(userData.Email)).
		Only(c.Context())
	if err != nil {
		return err
	}

	valid, err := argon2.VerifyEncoded([]byte(userData.Password), []byte(u.Password))
	if err != nil {
		return err
	}

	if !valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid password"})
	}

	userAgent, ok := c.GetReqHeaders()["User-Agent"]
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "User-Agent header not found"})
	}

	s, err := database.DatabaseClient.Session.
		Create().
		SetID(uuid.New()).
		SetUserID(u.ID).
		SetIP(c.IP()).
		SetUserAgent(userAgent).
		Save(c.Context())

	if err != nil {
		return err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  s.ID.String(),
		"type": "session",
	})

	tokenString, err := token.SignedString([]byte(config.Environment.JwtKey))

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": tokenString,
	})
}
