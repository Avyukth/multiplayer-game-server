package handlers_test

import (
	"context"
	"testing"

	gamestats "github.com/Avyukth/lila-assgnm/api/proto"
	"github.com/Avyukth/lila-assgnm/connections"
	"github.com/Avyukth/lila-assgnm/handlers"
	"github.com/go-redis/redismock/v9"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MockedCollection struct {
	*mongo.Database
}

func (m *MockedCollection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	// Mocked function
	return nil, nil
}

func TestGetGameStats(t *testing.T) {
	db, mock := redismock.NewClientMock()
	mock.ExpectGet("areaCode:1").SetVal("1")
	mongoClient := &MockedCollection{}
	server := handlers.Server{
		Mongo: &connections.MongoConnection{Db: mongoClient.Database},
		Redis: &connections.RedisConnection{Client: db},
	}

	req := &gamestats.GameStatsRequest{AreaCode: 1}
	resp, err := server.GetGameStats(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, req.AreaCode, resp.AreaCode)
	// You can add more assertions depending on what you expect from your function
}
