package user

import (
	"errors"
	"image"
	"io"
	"path"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/fyralabs/id-server/config"
	"github.com/fyralabs/id-server/ent"
	"github.com/fyralabs/id-server/util"
	"github.com/gofiber/fiber/v2"

	// "golang.org/x/image/webp"
	_ "image/jpeg"
	_ "image/png"

	"github.com/chai2010/webp"
	_ "golang.org/x/image/webp"
	"golang.org/x/sync/errgroup"
)

func UploadAvatar(c *fiber.Ctx) error {
	user := c.Locals("user").(*ent.User)

	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Request is not multipart/form-data",
		})
	}

	val, ok := form.File["avatar"]
	if !ok || len(val) == 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "No file uploaded",
		})
	}

	file, err := val[0].Open()
	if err != nil {
		return err
	}

	defer file.Close()

	src, _, err := image.Decode(file)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid Image",
		})
	}

	type SubImager interface {
		SubImage(r image.Rectangle) image.Image
	}

	subImager, ok := src.(SubImager)
	if !ok {
		return errors.New("image does not implement SubImage")
	}

	startX := 0
	startY := 0

	if src.Bounds().Dx() > 500 {
		startX = (src.Bounds().Dx() - 500) / 2
	}

	if src.Bounds().Dy() > 500 {
		startY = (src.Bounds().Dy() - 500) / 2
	}

	dst := subImager.SubImage(image.Rect(startX, startY, 500+startX, 500+startY))

	group, _ := errgroup.WithContext(c.Context())
	reader, writer := io.Pipe()
	objectPath := path.Join(user.ID.String(), "avatar.webp")

	group.Go(func() error {
		defer writer.Close()
		if err := webp.Encode(writer, dst, &webp.Options{
			Lossless: true,
		}); err != nil {
			return err
		}

		return nil
	})

	group.Go(func() error {
		defer reader.Close()
		if _, err = util.UploadClient.Upload(
			&s3manager.UploadInput{
				Bucket:      aws.String(config.Environment.S3Bucket),
				Key:         aws.String(objectPath),
				ContentType: aws.String("image/webp"),
				Body:        reader,
			},
		); err != nil {
			return err
		}

		return nil
	})

	if err := group.Wait(); err != nil {
		return err
	}

	url := "https://accounts-cdn.fyralabs.com/" + objectPath

	if _, err := user.Update().SetAvatarURL(url).Save(c.Context()); err != nil {
		return err
	}

	return c.SendStatus(200)
}

func DeleteAvatar(c *fiber.Ctx) error {
	user := c.Locals("user").(*ent.User)

	if user.AvatarURL == nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "No avatar to delete",
		})
	}

	objectPath := path.Join(user.ID.String(), "avatar.webp")

	if _, err := util.S3Client.DeleteObjectWithContext(
		c.Context(),
		&s3.DeleteObjectInput{
			Bucket: aws.String(config.Environment.S3Bucket),
			Key:    aws.String(objectPath),
		},
	); err != nil {
		return err
	}

	if _, err := user.Update().ClearAvatarURL().Save(c.Context()); err != nil {
		return err
	}

	return c.SendStatus(200)
}
