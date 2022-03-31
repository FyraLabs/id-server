package util

import (
	"github.com/fyralabs/id-server/config"
	"github.com/fyralabs/id-server/ent"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"os"
)

var SendGridClient *sendgrid.Client

func InitializeSendGrid() {
	SendGridClient = sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
}

func GenerateEmailVerificationToken(userID uuid.UUID, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   userID.ID(),
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

	to := mail.NewEmail(user.Name, user.Email)
	from := mail.NewEmail("FiraLabs ID", "noreply@fyralabs.com")
	subject := "Verify Your FyraLabs ID"
	plainTextContent := "Hey" + user.Name + ",\n" + "Welcome to FyraLabs ID. Please click on the link below to verify your email!\n" + "https://id.fyralabs.com/verifyEmail?token=" + tokenString
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, "")

	if _, err := SendGridClient.Send(message); err != nil {
		return err
	}

	return nil
}
