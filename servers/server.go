package main

import (
	"context"
	"log"
	"net"
	"os"
	"time"

	gamestats "github.com/Avyukth/lila-assgnm/api/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"gopkg.in/mgo.v2/bson"
)

// var db *mongo.Database
type server struct {
	gamestats.UnimplementedGameStatsServer
	Mongo *MongoConnection
}

type gameModes struct {
	BattleRoyal     int32 `bson:"battle_royal"`
	TeamDeathMatch  int32 `bson:"team_death_match"`
	CaptureTheFlag  int32 `bson:"capture_the_flag"`
}

type gameStat struct {
	AreaCode  int32     `bson:"area_code"`
	GameModes gameModes `bson:"game_modes"`
}


func (s *server) GetGameStats(ctx context.Context, in *gamestats.GameStatsRequest) (*gamestats.GameStatsResponse, error) {
	filter := bson.M{"area_code": in.AreaCode}
	gameStatsCollection := s.Mongo.Db.Collection("gameStats")
	cursor, err := gameStatsCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var gameModes []*gamestats.GameMode
	for cursor.Next(ctx) {
		var result gameStat
		err := cursor.Decode(&result)
		if err != nil {
			return nil, err
		}

		gameModes = append(gameModes,
			&gamestats.GameMode{Mode: "Battle Royale", Players: result.GameModes.BattleRoyal},
			&gamestats.GameMode{Mode: "Team Death Match", Players: result.GameModes.TeamDeathMatch},
			&gamestats.GameMode{Mode: "Capture The Flag", Players: result.GameModes.CaptureTheFlag},
		)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &gamestats.GameStatsResponse{
		AreaCode:  in.AreaCode,
		GameModes: gameModes,
	}, nil
}



type MongoConnection struct {
	Client *mongo.Client
	Db     *mongo.Database
}


type MongoConfig struct {
	Uri        string
	Database   string
	MaxPoolSize uint64
}

func NewMongoConnection(cfg MongoConfig) (*MongoConnection, error) {
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

// Close closes the MongoDB connection.
func (mc *MongoConnection) Close() {
	err := mc.Client.Disconnect(context.Background())
	if err != nil {
		log.Println("Error disconnecting from MongoDB:", err)
	}
}

func main() {
	// connect to MongoDB
	cfg := MongoConfig{
		Uri:        "mongodb://root:example@localhost:27017/?directConnection=true",
		Database:   "gameDB",
		MaxPoolSize: 50,
	}

	mongoConn, err := NewMongoConnection(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	defer mongoConn.Close()
	// start the gRPC server
	grpcServer := grpc.NewServer()
	gamestats.RegisterGameStatsServer(grpcServer, &server{Mongo: mongoConn})

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
	os.Exit(0)
}
