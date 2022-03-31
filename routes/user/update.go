package user

import (
	"encoding/json"
	"github.com/fyralabs/id-server/ent"
	"github.com/gofiber/fiber/v2"
)

type UpdateUser struct {
	Email *string `json:"email" validate:"omitempty,email,min=5,max=256"`
	Name  *string `json:"name" validate:"omitempty,min=1,max=256"`
}

func UpdateMe(c *fiber.Ctx) error {
	body := c.Request().Body()

	var updateData UpdateUser
	if err := json.Unmarshal(body, &updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid JSON"})
	}

	if err := validate.Struct(updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	user := c.Locals("user").(*ent.User)

	u := user.Update()

	if updateData.Email != nil {
		u = u.SetEmail(*updateData.Email)
	}

	if updateData.Name != nil {
		u = u.SetName(*updateData.Name)
	}

	if _, err := u.Save(c.Context()); err != nil {
		return err
	}

	return c.Status(200).Send(nil)
}