package database

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestElasticConnection(t *testing.T) {
	elasticUrl := "http://localhost:9205"
	client, err := elastic.NewClient(elastic.SetURL(elasticUrl))

	if err != nil {
		t.Fatalf("Failed to create elastic connection: %v", err)
	}

	info, code, err := client.Ping(elasticUrl).Do(context.Background())
	if err != nil {
		t.Fatalf("Failed to ping elastic connection: %v", err)
	}

	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	assert.NoError(t, err, "Failed to create elastic connection")
}
