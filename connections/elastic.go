package connections

import (
	"log"
	"os"

	"github.com/olivere/elastic/v7"
)

var ElasticConn *ElasticConnection

type ElasticConnection struct {
	Client *elastic.Client
}

type elasticConfig struct {
	Address  string
	Username string
	Password string
}

func newElasticConnection(cfg elasticConfig) (*ElasticConnection, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(cfg.Address),
		elastic.SetBasicAuth(cfg.Username, cfg.Password),
	)

	if err != nil {
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
		Address:  os.Getenv("ELASTIC_ADDRESS"),
		Username: os.Getenv("ELASTIC_USERNAME"),
		Password: os.Getenv("ELASTIC_PASSWORD"),
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
