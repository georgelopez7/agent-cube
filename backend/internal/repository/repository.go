package repository

import "agent-cube/internal/pkg/mongo"

type Repository struct {
	mongo *mongo.Mongo
}

func NewRepository(mongo *mongo.Mongo) *Repository {
	return &Repository{
		mongo: mongo,
	}
}
