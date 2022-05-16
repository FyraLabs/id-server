package util

import "github.com/gofiber/fiber/v2"

func GetClientIP(c *fiber.Ctx) string {
	if c.IsProxyTrusted() {
		ips := c.IPs()
		if len(ips) > 0 {
			return ips[0]
		}
	}

	return c.IP()
}
