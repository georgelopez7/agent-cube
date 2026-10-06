package test

import (
	"os"
	"testing"

	"agent-cube/internal/pkg/mongo"
	"agent-cube/internal/repository"
)

var repo *repository.Repository

func TestMain(m *testing.M) {
	uri, teardown := mongo.SetupMockMongo()
	defer teardown()

	mongoDB := mongo.NewMongoDB(uri, "test-db")

	repo = repository.NewRepository(mongoDB)

	os.Exit(m.Run())
}
