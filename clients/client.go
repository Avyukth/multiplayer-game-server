package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	gamestats "github.com/Avyukth/lila-assgnm/api/proto"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func mockClientRequest(id int, client gamestats.GameStatsClient, areaCode int32) {
	req := &gamestats.GameStatsRequest{
		AreaCode: areaCode,
	}

	res, err := client.GetGameStats(context.Background(), req)
	if err != nil {
		log.Fatalf("Client %d: Error invoking GetGameStats: %v", id, err)
	}

	log.Printf("Client %d: Received response for area code %d:\n", id, areaCode)
	for _, mode := range res.GetGameModes() {
		log.Printf("Mode: %s, Players: %d\n", mode.GetMode(), mode.GetPlayers())
	}
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

func main() {
	server := os.Getenv("SERVER")                // Assuming the .env file has a SERVER variable
	port, err := strconv.Atoi(os.Getenv("PORT")) // Assuming the .env file has a PORT variable

	if err != nil {
		log.Fatalf("Invalid PORT value: %v", err)
	}

	address := server + ":" + strconv.Itoa(port)

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	client := gamestats.NewGameStatsClient(conn)

	var wg sync.WaitGroup
	areaCodes := []int32{183, 184, 185, 123, 467, 789, 981}
	for i, areaCode := range areaCodes {
		wg.Add(1)
		go func(i int, areaCode int32) {
			defer wg.Done()
			mockClientRequest(i+1, client, areaCode)
		}(i, areaCode)
	}

	wg.Wait()
}
