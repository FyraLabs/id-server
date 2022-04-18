package user

import (
	"encoding/json"
	"time"

	"github.com/fyralabs/id-server/database"
	"github.com/fyralabs/id-server/ent"
	"github.com/fyralabs/id-server/ent/totpmethod"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/pquerna/otp/totp"
	"github.com/samber/lo"
)

type AddMethodData struct {
	Method string          `json:"method" validate:"required,oneof='totp'"`
	Name   string          `json:"name" validate:"required,min=1,max=256"`
	Data   json.RawMessage `json:"data" validate:"required"`
}

type TOTPData struct {
	// TODO: Maybe check if this is base32
	Secret string `json:"secret" validate:"required"`
	Code   string `json:"code" validate:"required,len=6,number"`
}

// var base32 = regexp.MustCompile("^[A-Z2-7]*$")
type GenericMethodResponse struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	Type       string     `json:"type"`
	CreatedAt  time.Time  `json:"createdAt"`
	LastUsedAt *time.Time `json:"lastUsedAt"`
}

func GetMethods(c *fiber.Ctx) error {
	user := c.Locals("user").(*ent.User)

	totpMethods, err := user.QueryTotpMethods().All(c.Context())

	if err != nil {
		return err
	}

	res := lo.Map(totpMethods, func(s *ent.TOTPMethod, _ int) GenericMethodResponse {
		return GenericMethodResponse{
			ID:         s.ID.String(),
			Type:       "totp",
			Name:       s.Name,
			CreatedAt:  s.CreatedAt,
			LastUsedAt: s.LastUsedAt,
		}
	})

	return c.JSON(res)
}

func RemoveMethod(c *fiber.Ctx) error {
	user := c.Locals("user").(*ent.User)

	methodIdString := c.Params("id")

	methodId, err := uuid.Parse(methodIdString)
	if err != nil {
		return err
	}

	existsTotp, err := user.QueryTotpMethods().Where(totpmethod.ID(methodId)).Exist(c.Context())
	if err != nil {
		return err
	}

	if existsTotp {
		if err := database.DatabaseClient.TOTPMethod.DeleteOneID(methodId).Exec(c.Context()); err != nil {
			return err
		}

		return c.SendStatus(fiber.StatusOK)
	}

	return c.SendStatus(fiber.StatusNotFound)
}

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
			if err := json.Unmarshal(methodData.Data, &methodParams); err != nil {
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
				SetID(uuid.New()).
				SetUserID(user.ID).
				SetName(methodData.Name).
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
