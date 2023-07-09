package main

import (
	"context"
	"log"
	"net"
	"os"

	gamestats "github.com/Avyukth/lila-assgnm/api/proto"
	"google.golang.org/grpc"
)



type server struct {
	gamestats.UnimplementedGameStatsServer
}

func (s *server) GetGameStats(ctx context.Context, in *gamestats.GameStatsRequest) (*gamestats.GameStatsResponse, error) {
	// Static data for example
	return &gamestats.GameStatsResponse{
		AreaCode: in.AreaCode, // Echo back the requested area code
		GameModes: []*gamestats.GameMode{
			&gamestats.GameMode{Mode: "Battle Royale", Players: 10},
			&gamestats.GameMode{Mode: "Team Death Match", Players: 20},
			&gamestats.GameMode{Mode: "Capture the Flag", Players: 30},
		},
	}, nil
}


func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	gamestats.RegisterGameStatsServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	os.Exit(0)
}
