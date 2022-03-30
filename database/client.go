package database

import (
	"context"
	"log"

	"github.com/fyralabs/id-server/ent"
	_ "github.com/lib/pq"
)

var DatabaseClient *ent.Client

func InitializeDatabase() error {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable	")
	if err != nil {
		return err
	}
	// defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	DatabaseClient = client

	return nil
}
