package main

import (
	"log"
	"net"
	"os"
	"strings"

	gamestats "github.com/Avyukth/lila-assgnm/api/proto"
	"github.com/Avyukth/lila-assgnm/connections"
	"github.com/Avyukth/lila-assgnm/handlers"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

type serverConfig struct {
	Port string
}

func init() {
	var envFilePath string
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			if strings.HasPrefix(arg, "-e=") {
				envFilePath = strings.TrimPrefix(arg, "-e=")
				break
			}
		}
	}

	if envFilePath == "" {
		envFilePath = "../.env"
	}

	// loads values from .env into the system
	if err := godotenv.Load(envFilePath); err != nil {
		log.Fatalf("No .env file found at path: %s", envFilePath)
	}
}

func loadServerConfig() serverConfig {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Failed to load PORT environment variable")
	}
	return serverConfig{
		Port: ":" + port, // prepend ":" to port number
	}
}

func main() {
	connections.InitElastic()
	defer connections.ElasticConn.Close()

	logMessage := connections.LogMessage{
		Level:   "INFO",
		Message: "This is a test log message",
	}

	connections.ElasticConn.SendLog("my-log-index", logMessage)

	connections.InitMongo()
	defer connections.MongoConn.Close()
	connections.InitRedis()
	defer connections.RedisConn.Close()

	server := loadServerConfig()

	grpcServer := grpc.NewServer()
	gamestats.RegisterGameStatsServer(grpcServer, &handlers.Server{Mongo: connections.MongoConn, Redis: connections.RedisConn})

	listener, err := net.Listen("tcp", server.Port)
	if err != nil {
		logMessage := connections.LogMessage{
			Level:   "FATAL",
			Message: "Failed to listen: " + err.Error(),
		}
		connections.ElasticConn.SendLog("my-log-index", logMessage)
		log.Fatalf("Failed to listen: %v", err)
	}

	logMessage = connections.LogMessage{
		Level:   "INFO",
		Message: "Server is running on gRPC server with port " + server.Port + " ...",
	}
	connections.ElasticConn.SendLog("my-log-index", logMessage)
	log.Println("Server is running on gRPC server with port " + server.Port + " ...")

	if err := grpcServer.Serve(listener); err != nil {
		logMessage := connections.LogMessage{
			Level:   "FATAL",
			Message: "Failed to serve: " + err.Error(),
		}
		connections.ElasticConn.SendLog("my-log-index", logMessage)
		log.Fatalf("Failed to serve: %v", err)
	}

	os.Exit(0)
}
