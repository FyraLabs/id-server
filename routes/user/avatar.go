package user

import (
	"image"
	"io"
	"path"

	"github.com/fyralabs/id-server/config"
	"github.com/fyralabs/id-server/ent"
	"github.com/fyralabs/id-server/util"
	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
	"golang.org/x/image/draw"

	// "golang.org/x/image/webp"
	"github.com/chai2010/webp"
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

	dst := image.NewRGBA(image.Rect(0, 0, 500, 500))
	draw.NearestNeighbor.Scale(dst, dst.Rect, src, src.Bounds(), draw.Over, nil)

	reader, writer := io.Pipe()

	if err := webp.Encode(writer, dst, &webp.Options{}); err != nil {
		return err
	}

	objectPath := path.Join(user.ID.String(), "avatar.webp")

	if _, err = util.S3Client.PutObject(
		c.Context(),
		config.Environment.S3AvatarBucket,
		objectPath,
		reader,
		// TODO: Get the size of the file, this wastes memory
		-1,
		minio.PutObjectOptions{ContentType: "image/webp"},
	); err != nil {
		return err
	}

	url := config.Environment.S3AvatarURLPrefix + objectPath

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

	if err := util.S3Client.RemoveObject(
		c.Context(),
		config.Environment.S3AvatarBucket,
		objectPath,
		minio.RemoveObjectOptions{},
	); err != nil {
		return err
	}

	if _, err := user.Update().ClearAvatarURL().Save(c.Context()); err != nil {
		return err
	}

	return c.SendStatus(200)
}
