package user

import (
	"encoding/json"

	"github.com/fyralabs/id-server/database"
	"github.com/fyralabs/id-server/ent"
	"github.com/gofiber/fiber/v2"
	"github.com/pquerna/otp/totp"
)

type AddMethodData struct {
	Method string          `json:"method" validate:"required,oneof='totp'"`
	Data   json.RawMessage `json:"data" validate:"required"`
}

type TOTPData struct {
	// TODO: Maybe check if this is base32
	Secret string `json:"secret" validate:"required"`
	Code   string `json:"method" validate:"required,len=6,number"`
}

// var base32 = regexp.MustCompile("^[A-Z2-7]*$")

func AddMethod(c *fiber.Ctx) error {
	user := c.Locals("user").(*ent.User)
	body := c.Request().Body()

	var methodData AddMethodData
	if err := json.Unmarshal(body, &methodData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid JSON"})
	}

	if err := validate.Struct(methodData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	switch methodData.Method {
	case "totp":
		{
			var methodParams TOTPData
			if err := json.Unmarshal(body, &methodParams); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid JSON"})
			}

			if err := validate.Struct(methodParams); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
			}

			if ok := totp.Validate(methodParams.Code, methodParams.Secret); !ok {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid code"})
			}

			m, err := database.DatabaseClient.TOTPMethod.
				Create().
				SetUserID(user.ID).
				SetSecret(methodParams.Secret).
				Save(c.Context())

			if err != nil {
				return err
			}

			return c.Status(fiber.StatusCreated).JSON(fiber.Map{
				"id": m.ID,
			})
		}
	}

	return c.SendStatus(fiber.StatusBadRequest)
}
