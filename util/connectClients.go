package util

import (
	"strings"

	"github.com/fyralabs/id-server/config"
)

var ConnectClientIDs = map[string]string{}

func InitializeConnectClients() {
	ids := strings.Split(config.Environment.ClientConnectIDs, ",")

	for _, id := range ids {
		ConnectClientIDs[id] = id
	}
}
