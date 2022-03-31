package database

import (
	"context"
	"github.com/fyralabs/id-server/config"
	"log"

	"github.com/fyralabs/id-server/ent"
	_ "github.com/lib/pq"
)

var DatabaseClient *ent.Client

func InitializeDatabase() error {
	client, err := ent.Open("postgres", config.Environment.DatabaseOptions)
	if err != nil {
		return err
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	DatabaseClient = client

	return nil
}
