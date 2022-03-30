package user

import (
	"encoding/json"

	"github.com/fyralabs/id-server/database"
	"github.com/fyralabs/id-server/ent/user"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/matthewhartstonge/argon2"
)

var validate = validator.New()

type RegisterUser struct {
	Name     string `json:"name" validate:"required,min=1,max=256"`
	Email    string `json:"email" validate:"required,email,min=5,max=256"`
	Password string `json:"password" validate:"required,min=8,max=256"`
}

func Register(c *fiber.Ctx) error {
	body := c.Request().Body()

	var userData RegisterUser
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

	if exists {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"message": "User already exists"})
	}

	argon := argon2.DefaultConfig()

	encoded, err := argon.HashEncoded([]byte(userData.Password))
	if err != nil {
		return err
	}

	u, err := database.DatabaseClient.User.
		Create().
		SetID(uuid.New()).
		SetName(userData.Name).
		SetEmail(userData.Email).
		SetPassword(string(encoded)).
		Save(c.Context())

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id": u.ID,
	})
}
