package meilidb

import (
	"os"

	"github.com/meilisearch/meilisearch-go"
)

// GetClient reutun the client
func GetClient() (client *meilisearch.Client) {

	host := os.Getenv("MEILI_HOST")
	apiKey := os.Getenv("MEILI_KEY")
	client = meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   host,
		APIKey: apiKey,
	})

	return
}
