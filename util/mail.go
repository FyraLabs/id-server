package util

import (
	"bytes"
	"embed"
	"errors"

	"html/template"

	"github.com/fyralabs/id-server/config"
	"github.com/fyralabs/id-server/ent"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

//go:embed emails/*
var emails embed.FS
var SendGridClient *sendgrid.Client

func InitializeSendGrid() {
	SendGridClient = sendgrid.NewSendClient(config.Environment.SendGridKey)
}

func GenerateEmailVerificationToken(userID uuid.UUID, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   userID.String(),
		"email": email,
		"type":  "emailVerification",
	})

	tokenString, err := token.SignedString([]byte(config.Environment.JwtKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func SendVerificationEmail(user *ent.User) error {
	tokenString, err := GenerateEmailVerificationToken(user.ID, user.Email)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)

	a, err := template.ParseFS(emails, "emails/verify.html")

	if err != nil {
		return err
	}

	data := struct {
		ConfirmURL string
	}{
		ConfirmURL: "https://accounts.fyralabs.com/verifyEmail?token=" + tokenString,
	}

	a.Execute(buf, data)

	to := mail.NewEmail(user.Name, user.Email)
	from := mail.NewEmail("Fyra Accounts", "noreply@fyralabs.com")
	subject := "Verify Your Fyra Account"
	message := mail.NewSingleEmail(from, subject, to, "", buf.String())

	r, err := SendGridClient.Send(message)
	if err != nil {
		return err
	}

	if r.StatusCode != 202 {
		return errors.New(r.Body)
	}

	return nil
}
