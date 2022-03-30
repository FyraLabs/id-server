package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Auth(c *fiber.Ctx) error {
	tokenString, ok := c.GetRespHeaders()["Authorization"]

	if !ok {
		return c.Status(401).JSON(fiber.Map{"message": "Unauthorized"})
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return hmacSampleSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		fmt.Println(err)
	}

	return c.Next()
}
