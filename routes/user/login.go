package user

import (
	"encoding/json"
	"time"

	"github.com/fyralabs/id-server/config"
	"github.com/fyralabs/id-server/ent"
	"github.com/fyralabs/id-server/util"
	"github.com/pquerna/otp/totp"
	"github.com/samber/lo"

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

	u, err := database.DatabaseClient.User.
		Query().
		Where(user.EmailEQ(userData.Email)).
		Only(c.Context())

	if err != nil {
		if err.Error() == "ent: user not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User doesn't exist"})
		}

		return err
	}

	valid, err := argon2.VerifyEncoded([]byte(userData.Password), []byte(u.Password))
	if err != nil {
		return err
	}

	if !valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid password"})
	}

	// TODO: Check for other methods
	totpMethods, err := u.QueryTotpMethods().All(c.Context())
	if err != nil {
		return err
	}

	if len(totpMethods) > 0 {
		res := lo.Map(totpMethods, func(s *ent.TOTPMethod, _ int) fiber.Map {
			return fiber.Map{
				"id":   s.ID.String(),
				"type": "totp",
				"name": s.Name,
			}
		})

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub":  u.ID,
			"type": "2fa",
			"exp":  jwt.NewNumericDate(time.Now().Add(time.Minute * 15)),
		})

		tokenString, err := token.SignedString([]byte(config.Environment.JwtKey))

		if err != nil {
			return err
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"type": "2fa",
			"data": fiber.Map{
				"methods": res,
				"token":   tokenString,
			},
		})
	}

	userAgent, ok := c.GetReqHeaders()["User-Agent"]
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "User-Agent header not found"})
	}

	tokenString, err := util.CreateSession(u.ID, userAgent, c.IP(), c.Context())
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"type": "session",
		"data": fiber.Map{
			"token": tokenString,
		},
	})
}

type Login2FAParams struct {
	Token    string          `json:"token" validate:"required,jwt"`
	Method   string          `json:"method" validate:"required,oneof='totp'"`
	MethodID string          `json:"method_id" validate:"required,uuid"`
	Data     json.RawMessage `json:"data" validate:"required"`
}

type LoginTOTP struct {
	Code string `json:"code" validate:"required,len=6,number"`
}

func Login2FA(c *fiber.Ctx) error {
	body := c.Request().Body()

	var reqData Login2FAParams
	if err := json.Unmarshal(body, &reqData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid JSON"})
	}

	if err := validate.Struct(reqData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	claims, err := util.DecodeJWT(reqData.Token, "2fa")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid token"})
	}

	userIDString, ok := claims["sub"].(string)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"message": "Invalid token"})
	}

	parse, err := uuid.Parse(userIDString)
	if err != nil {
		return err
	}

	u, err := database.DatabaseClient.User.
		Query().
		Where(user.IDEQ(parse)).
		Only(c.Context())

	if err != nil {
		if err.Error() == "ent: user not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User doesn't exist"})
		}

		return err
	}

	switch reqData.Method {
	case "totp":
		{
			var methodParams LoginTOTP
			if err := json.Unmarshal(reqData.Data, &methodParams); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid JSON"})
			}

			if err := validate.Struct(methodParams); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
			}

			totpMethods, err := u.QueryTotpMethods().All(c.Context())

			if err != nil {
				return err
			}

			totpMethod, ok := lo.Find(totpMethods, func(s *ent.TOTPMethod) bool {
				return s.ID.String() == reqData.MethodID
			})

			if !ok {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "TOTP method not found"})
			}

			if ok := totp.Validate(methodParams.Code, totpMethod.Secret); !ok {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid code"})
			}

			_, err = totpMethod.Update().SetLastUsedAt(time.Now()).Save(c.Context())
			if err != nil {
				return err
			}

			userAgent, ok := c.GetReqHeaders()["User-Agent"]
			if !ok {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "User-Agent header not found"})
			}

			tokenString, err := util.CreateSession(u.ID, userAgent, c.IP(), c.Context())
			if err != nil {
				return err
			}

			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"token": tokenString,
			})
		}
	}

	return c.SendStatus(fiber.StatusBadRequest)
}
