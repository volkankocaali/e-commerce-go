package database

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/volkankocaali/e-commorce-go/pkg/config"
	"log"
)

func NewElasticClient(cfg config.Config) (*elastic.Client, error) {
	elasticUrl := fmt.Sprintf("http://%s:%s", cfg.ElasticSearchHost, cfg.ElasticSearchPort)
	client, err := elastic.NewClient(elastic.SetURL(elasticUrl), elastic.SetSniff(false))
	if err != nil {
		log.Fatal(err)
	}

	info, code, err := client.Ping(elasticUrl).Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	return client, nil
}
