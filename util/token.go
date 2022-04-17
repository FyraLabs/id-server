package util

import (
	"context"
	"errors"
	"fmt"

	"github.com/fyralabs/id-server/config"
	"github.com/fyralabs/id-server/database"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func DecodeJWT(tokenString string, tokenType string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.Environment.JwtKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return nil, errors.New("cannot parse token")
	}

	if err := claims.Valid(); err != nil {
		return nil, errors.New("invalid token")
	}

	if decodedType, ok := claims["type"].(string); !ok || decodedType != tokenType {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func CreateSession(userID uuid.UUID, userAgent string, ip string, context context.Context) (string, error) {
	s, err := database.DatabaseClient.Session.
		Create().
		SetID(uuid.New()).
		SetUserID(userID).
		SetIP(ip).
		SetUserAgent(userAgent).
		Save(context)

	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  s.ID.String(),
		"type": "session",
	})

	tokenString, err := token.SignedString([]byte(config.Environment.JwtKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
