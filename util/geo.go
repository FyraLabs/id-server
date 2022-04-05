package util

import (
	"github.com/oschwald/geoip2-golang"
)

var GeoIP *geoip2.Reader

func InitializeGeoIP() error {
	db, err := geoip2.Open("./GeoLite2-City.mmdb")
	if err != nil {
		return err
	}

	GeoIP = db

	return nil
}