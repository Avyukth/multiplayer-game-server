package connections

import (
	"context"
	"log"
	"os"

	"github.com/olivere/elastic/v7"
)

var ElasticConn *ElasticConnection

type LogMessage struct {
	Level   string `json:"level"`
	Message string `json:"message"`
}
type ElasticConnection struct {
	Client *elastic.Client
}

type elasticConfig struct {
	Address string
}

func (ec *ElasticConnection) SendLog(index string, logMessage LogMessage) {
	ctx := context.Background()

	_, err := ec.Client.Index().
		Index(index).
		BodyJson(logMessage).
		Do(ctx)
	if err != nil {
		// Handle error
		log.Println(err)
	}
}

func newElasticConnection(cfg elasticConfig) (*ElasticConnection, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(cfg.Address),
	)

	if err != nil {
		log.Println("Error creating Elasticsearch client: ", err)
		return nil, err
	}

	log.Println("Connected to Elasticsearch")

	return &ElasticConnection{Client: client}, nil
}

func (ec *ElasticConnection) Close() {
	ec.Client.Stop()

	log.Println("Failed to close Elasticsearch connection")

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
