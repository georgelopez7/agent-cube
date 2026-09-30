package mongo

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Collection string

const (
	RubiksCubes Collection = "rubiks_cubes"
)

type Mongo struct {
	Client *mongo.Client
	DB     *mongo.Database
}

func NewMongoDB(uri string, dbName string) *Mongo {
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalln(err)
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		log.Fatal("failed to ping mongo: ", err)
	}

	return &Mongo{
		Client: client,
		DB:     client.Database(dbName),
	}
}

func (m *Mongo) GetCollection(name Collection) *mongo.Collection {
	return m.DB.Collection(string(name))
}
