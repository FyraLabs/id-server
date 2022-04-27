package util

import (
	client "github.com/ory/hydra-client-go"
)

var HydraAdmin *client.APIClient

func InitializeHydra() {
	configuration := client.NewConfiguration()
	configuration.Servers = []client.ServerConfiguration{
		{
			URL: "http://localhost:4445", // Admin API URL
		},
	}

	HydraAdmin = client.NewAPIClient(configuration)
}
