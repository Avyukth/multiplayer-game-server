package connections

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

var ElasticConn *ElasticConnection

type LogMessage struct {
	Level   string `json:"level"`
	Message string `json:"message"`
}

type ElasticConnection struct {
	Client *elasticsearch.Client
}

type elasticConfig struct {
	Address string
}

func (ec *ElasticConnection) SendLog(index string, logMessage LogMessage) {
	ctx := context.Background()

	// Marshal the LogMessage to JSON
	jsonBody, err := json.Marshal(logMessage)
	if err != nil {
		log.Println("Error marshaling log message: ", err)
		return
	}

	req := esapi.IndexRequest{
		Index:   index,
		Body:    strings.NewReader(string(jsonBody)),
		Refresh: "true",
	}

	res, err := req.Do(ctx, ec.Client)
	if err != nil {
		// Handle error
		log.Println(err)
	}

	if res.IsError() {
		log.Println(res.String())
	}
}

func newElasticConnection(cfg elasticConfig) (*ElasticConnection, error) {
	config := elasticsearch.Config{
		Addresses: []string{
			cfg.Address,
		},
	}

	client, err := elasticsearch.NewClient(config)
	if err != nil {
		log.Println("Error creating Elasticsearch client: ", err)
		return nil, err
	}

	log.Println("Connected to Elasticsearch")

	return &ElasticConnection{Client: client}, nil
}

func (ec *ElasticConnection) Close() {
	// No explicit stop method in the official Elasticsearch Go client
	log.Println("Closing Elasticsearch connection")
}

func LoadElasticConfig() elasticConfig {
	elasticConfig := elasticConfig{
		Address: os.Getenv("ELASTIC_ADDRESS"),
	}
	return elasticConfig
}

func InitElastic() {
	var err error
	ElasticConn, err = newElasticConnection(LoadElasticConfig())
	if err != nil {
		log.Fatalf("Failed to connect to Elasticsearch: %v", err)
	}
}
