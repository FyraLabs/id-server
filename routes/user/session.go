package user

import (
	"net"

	"github.com/fyralabs/id-server/database"
	"github.com/fyralabs/id-server/ent"
	"github.com/fyralabs/id-server/ent/session"
	"github.com/fyralabs/id-server/util"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	ua "github.com/mileusna/useragent"
	"github.com/samber/lo"
)

func GetSessions(c *fiber.Ctx) error {
	user := c.Locals("user").(*ent.User)

	sessions, err := user.QuerySessions().All(c.Context())
	if err != nil {
		return err
	}

	res, err := lo.MapWithError(sessions, func(s *ent.Session, _ int) (fiber.Map, error) {
		ip := net.ParseIP(s.IP)
		record, err := util.GeoIP.City(ip)
		if err != nil {
			return nil, err
		}

		country := record.Country.Names["en"]
		subdivision := ""
		city := record.City.Names["en"]

		if len(record.Subdivisions) > 0 {
			subdivision = record.Subdivisions[0].Names["en"]
		}
		
		rec := fiber.Map{
			"id":         s.ID.String(),
			"ip":         s.IP,
			"userAgent":  s.UserAgent,
			"createdAt":  s.CreatedAt,
			"lastUsedAt": s.LastUsedAt,
		}
		
		if country != "" {
			rec["country"] = country
		}
		
		if city != "" {
			rec["city"] = city
		}
		
		if subdivision != "" {
			rec["subdivision"] = subdivision
		}
		
		ua := ua.Parse(s.UserAgent)

		if ua.Desktop {
			rec["device"] = "desktop"
		} else if ua.Mobile {
			rec["device"] = "mobile"
		} else if ua.Tablet {
			rec["device"] = "tablet"
		}

		if ua.Name != "" {
			rec["uaName"] = ua.Name
		}

		if ua.Version != "" {
			rec["uaVersion"] = ua.Version
		}

		if ua.OS != "" {
			rec["osName"] = ua.OS
		}

		if ua.OSVersion != "" {
			rec["osVersion"] = ua.OSVersion
		}

		return rec, nil
	})

	if err != nil {
		return err
	}

	return c.JSON(res)
}

func RevokeSession(c *fiber.Ctx) error {
	user := c.Locals("user").(*ent.User)

	sessionIdString := c.Params("id")

	sessionId, err := uuid.Parse(sessionIdString)
	if err != nil {
		return err
	}

	// This should prevent someone from revoking someone else's session
	session, err := user.QuerySessions().Where(session.ID(sessionId)).Only(c.Context())
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "Session not found"})
	}

	if err := database.DatabaseClient.Session.DeleteOne(session).Exec(c.Context()); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}
