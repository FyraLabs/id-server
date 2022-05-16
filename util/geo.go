package util

import (
	"github.com/fyralabs/id-server/config"
	"github.com/oschwald/geoip2-golang"
)

var GeoIP *geoip2.Reader

// TODO: we need to get the latest gepip and put it in docker
func InitializeGeoIP() error {
	db, err := geoip2.Open(config.Environment.GeoLite2CityPath)
	if err != nil {
		return err
	}

	GeoIP = db

	return nil
}
