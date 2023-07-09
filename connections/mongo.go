package connections

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoConn *MongoConnection
type MongoConnection struct {
	Client *mongo.Client
	Db     *mongo.Database
}

type mongoConfig struct {
	Uri        string
	Database   string
	MaxPoolSize uint64
}


func newMongoConnection(cfg mongoConfig) (*MongoConnection, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(cfg.Uri)
	clientOptions.SetMaxPoolSize(cfg.MaxPoolSize)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to MongoDB")

	db := client.Database(cfg.Database)
	return &MongoConnection{Client: client, Db: db}, nil
}

func (mc *MongoConnection) Close() {
	err := mc.Client.Disconnect(context.Background())
	if err != nil {
		log.Println("Error disconnecting from MongoDB:", err)
	}
}

func loadMongoConfig() mongoConfig {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("MONGODB_URI env var not set")
	}

    database := os.Getenv("MONGODB_DATABASE")
	if database == "" {
		log.Fatal("MONGODB_DATABASE env var not set")
	}

    maxPoolSize := os.Getenv("MONGODB_MAX_POOL_SIZE")
	if maxPoolSize == "" {
		log.Fatal("MONGODB_MAX_POOL_SIZE env var not set")
	}
	maxPoolSizeInt, err := strconv.ParseUint(maxPoolSize, 10, 64)
	if err != nil {
		log.Fatal("MONGODB_MAX_POOL_SIZE env var not set to a valid number")
	}

	mongoConfig := mongoConfig{
		Uri:        uri,
		Database:   database,
		MaxPoolSize: maxPoolSizeInt,
	}
	return mongoConfig
}


func InitMongo() {
	var err error
	MongoConn, err = newMongoConnection(loadMongoConfig())
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
}
