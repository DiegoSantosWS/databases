package es01db

import (
	"context"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

// GetConnect ...
func GetConnect(ctx context.Context) (newCtx context.Context) {
	newClient, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	})
	if err != nil {
		log.Panic("connect", err)
	}

	newCtx = context.WithValue(ctx, ClientKey, newClient)
	return
}
