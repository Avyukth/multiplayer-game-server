package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	gamestats "github.com/Avyukth/lila-assgnm/api/proto"
	"github.com/Avyukth/lila-assgnm/connections"
	"github.com/Avyukth/lila-assgnm/models"
	"github.com/redis/go-redis/v9"
	"gopkg.in/mgo.v2/bson"
)

type Server struct {
	gamestats.UnimplementedGameStatsServer
	Mongo *connections.MongoConnection
	Redis *connections.RedisConnection
}

func (s *Server) GetGameStats(ctx context.Context, in *gamestats.GameStatsRequest) (*gamestats.GameStatsResponse, error) {
	filter := bson.M{"area_code": in.AreaCode}
	areaCodeKey := fmt.Sprintf("areaCode:%d", in.AreaCode)

	cachedVal, err := s.Redis.Client.Get(ctx, areaCodeKey).Result()
	if err == redis.Nil {

		gameStatsCollection := s.Mongo.Db.Collection("gameStats")
		cursor, err := gameStatsCollection.Find(ctx, filter)
		if err != nil {
			log.Println("Error while retrieving from MongoDB:", err)
			return nil, err
		}
		defer cursor.Close(ctx)

		var gameModes []*gamestats.GameMode
		for cursor.Next(ctx) {
			var result models.GameStat
			err := cursor.Decode(&result)
			if err != nil {
				log.Println("Error while decoding data from MongoDB:", err)
				return nil, err
			}

			gameModes = append(gameModes,
				&gamestats.GameMode{Mode: "Battle Royale", Players: result.GameModes.BattleRoyal},
				&gamestats.GameMode{Mode: "Team Death Match", Players: result.GameModes.TeamDeathMatch},
				&gamestats.GameMode{Mode: "Capture The Flag", Players: result.GameModes.CaptureTheFlag},
			)
		}
		if err := cursor.Err(); err != nil {
			log.Println("Error after iterating over cursor:", err)
			return nil, err
		}

		data, err := json.Marshal(gameModes)
		if err != nil {
			log.Println("Error while marshaling data to JSON:", err)
			return nil, err
		}

		timeoutStr := os.Getenv("REDIS_TIMEOUT")
		timeout, err := strconv.Atoi(timeoutStr)
		if err != nil {
			log.Println("Error parsing REDIS_TIMEOUT, using default 5 minutes:", err)
			timeout = 5
		}

		err = s.Redis.Client.Set(ctx, areaCodeKey, data, time.Duration(timeout)*time.Minute).Err()
		if err != nil {
			log.Println("Error while setting data to Redis:", err)
			return nil, err
		}
		return &gamestats.GameStatsResponse{
			AreaCode:  in.AreaCode,
			GameModes: gameModes,
		}, nil
	} else if err != nil {
		log.Println("Error while getting data from Redis:", err)
		return nil, err
	}

	var gameModes []*gamestats.GameMode
	err = json.Unmarshal([]byte(cachedVal), &gameModes)
	if err != nil {
		log.Println("Error while unmarshaling data from Redis:", err)
		return nil, err
	}

	log.Println("Returning from cache: ", gameModes, "for area code:", in.AreaCode)

	return &gamestats.GameStatsResponse{
		AreaCode:  in.AreaCode,
		GameModes: gameModes,
	}, nil
}
