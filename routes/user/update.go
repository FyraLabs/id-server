package user

import (
	"encoding/json"

	"github.com/fyralabs/id-server/database"
	"github.com/fyralabs/id-server/ent"
	entUser "github.com/fyralabs/id-server/ent/user"
	"github.com/fyralabs/id-server/util"
	"github.com/gofiber/fiber/v2"
	"github.com/matthewhartstonge/argon2"
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
		exists, err := database.DatabaseClient.User.Query().Where(entUser.EmailEQ(*updateData.Email)).Exist(c.Context())
		if err != nil {
			return err
		}
	
		if exists {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"message": "User already exists"})
		}
		
		u = u.SetEmail(*updateData.Email).SetEmailVerified(false)
	}

	if updateData.Name != nil {
		u = u.SetName(*updateData.Name)
	}

	updatedUser, err := u.Save(c.Context())

	if err != nil {
		return err
	}

	if updateData.Email != nil {
		if err := util.SendVerificationEmail(updatedUser); err != nil {
			return err
		}
	}

	return c.SendStatus(fiber.StatusOK)
}

type UpdatePassword struct {
	CurrentPassword string `json:"currentPassword" validate:"required,min=8,max=256"`
	NewPassword  string `json:"newPassword" validate:"required,min=8,max=256"`
}

func UpdateMyPassword(c *fiber.Ctx) error {
	body := c.Request().Body()

	var updateData UpdatePassword
	if err := json.Unmarshal(body, &updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid JSON"})
	}

	if err := validate.Struct(updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	user := c.Locals("user").(*ent.User)

	valid, err := argon2.VerifyEncoded([]byte(updateData.CurrentPassword), []byte(user.Password))
	if err != nil {
		return err
	}

	if !valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid password"})
	}

	argon := argon2.DefaultConfig()

	encoded, err := argon.HashEncoded([]byte(updateData.NewPassword))
	if err != nil {
		return err
	}

	_, err = user.Update().SetPassword(string(encoded)).Save(c.Context())

	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}
