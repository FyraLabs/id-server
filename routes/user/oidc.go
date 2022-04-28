package user

import (
	"encoding/json"

	"github.com/fyralabs/id-server/ent"
	"github.com/fyralabs/id-server/util"
	"github.com/gofiber/fiber/v2"
	client "github.com/ory/hydra-client-go"
)

type OIDCLoginData struct {
	Challenge string `json:"challenge" validate:"required"`
}

func OIDCLogin(c *fiber.Ctx) error {
	user := c.Locals("user").(*ent.User)
	body := c.Request().Body()

	var challengeData OIDCLoginData
	if err := json.Unmarshal(body, &challengeData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid JSON"})
	}

	if err := validate.Struct(challengeData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	resp, _, err := util.HydraAdmin.AdminApi.AcceptLoginRequest(c.Context()).AcceptLoginRequest(client.AcceptLoginRequest{
		Subject: user.ID.String(),
	}).LoginChallenge(challengeData.Challenge).Execute()

	if err != nil {
		// TODO: Better error handling
		return c.Status(400).JSON(
			fiber.Map{
				"message": "Could not login with OIDC",
			},
		)
	}

	return c.Status(200).JSON(
		fiber.Map{
			"redirect": resp.GetRedirectTo(),
		},
	)
}

func OIDCLogout(c *fiber.Ctx) error {
	user := c.Locals("user").(*ent.User)

	resp, _, err := util.HydraAdmin.AdminApi.AcceptLogoutRequest(c.Context()).LogoutChallenge(user.ID.String()).Execute()

	if err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"message": "Could not logout with OIDC",
			},
		)
	}

	return c.Status(200).JSON(
		fiber.Map{
			"redirect": resp.GetRedirectTo(),
		},
	)
}
