package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"agent-cube/internal/domain"

	xmongo "agent-cube/internal/pkg/mongo"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// CreateRubiksCube - inserts a new rubiks cube into the collection.
func (r *Repository) CreateRubiksCube(ctx context.Context, cube domain.RubiksCube) error {
	collection := r.mongo.GetCollection(xmongo.RubiksCubes)

	_, err := collection.InsertOne(ctx, &cube)
	if err != nil {
		return fmt.Errorf("failed to create rubiks cube: %w", err)
	}

	return nil
}

// UpdateRubiksCube - replaces an existing rubiks cube by ID.
func (r *Repository) UpdateRubiksCube(ctx context.Context, cube *domain.RubiksCube) error {
	collection := r.mongo.GetCollection(xmongo.RubiksCubes)

	cube.UpdatedAt = time.Now().UTC()

	filter := bson.M{"_id": cube.ID}

	result, err := collection.ReplaceOne(ctx, filter, cube)
	if err != nil {
		return fmt.Errorf("failed to update rubiks cube: %w", err)
	}

	if result.MatchedCount == 0 {
		return errors.New("rubiks cube not found")
	}

	return nil
}

// UpdateRubiksCubeStatus - updates only the status (and updated_at) of a rubiks cube by ID.
func (r *Repository) UpdateRubiksCubeStatus(ctx context.Context, id primitive.ObjectID, status domain.RubiksCubeStatus) error {
	collection := r.mongo.GetCollection(xmongo.RubiksCubes)

	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"status":     status,
			"updated_at": time.Now().UTC(),
		},
	}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update rubiks cube status: %w", err)
	}

	if result.MatchedCount == 0 {
		return errors.New("rubiks cube not found")
	}

	return nil
}

// GetRubiksCubeByID - retrieves a rubiks cube by its ID.
func (r *Repository) GetRubiksCubeByID(ctx context.Context, id primitive.ObjectID) (*domain.RubiksCube, error) {
	collection := r.mongo.GetCollection(xmongo.RubiksCubes)

	filter := bson.M{"_id": id}

	var cube domain.RubiksCube
	err := collection.FindOne(ctx, filter).Decode(&cube)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get rubiks cube: %w", err)
	}

	return &cube, nil
}

// GetAllRubiksCubes - retrieves all rubiks cubes with an optional limit.
func (r *Repository) GetAllRubiksCubes(ctx context.Context, limit int64) ([]domain.RubiksCube, error) {
	collection := r.mongo.GetCollection(xmongo.RubiksCubes)

	opts := options.Find()
	if limit > 0 {
		opts.SetLimit(limit)
	}

	cursor, err := collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to list rubiks cubes: %w", err)
	}
	defer cursor.Close(ctx)

	var cubes []domain.RubiksCube
	if err := cursor.All(ctx, &cubes); err != nil {
		return nil, fmt.Errorf("failed to decode rubiks cubes: %w", err)
	}

	return cubes, nil
}

// ResetRubiksCubes - removes every document in the collection (used by tests).
func (r *Repository) ResetRubiksCubes(ctx context.Context) error {
	collection := r.mongo.GetCollection(xmongo.RubiksCubes)

	_, err := collection.DeleteMany(ctx, bson.M{})
	if err != nil {
		return fmt.Errorf("failed to reset rubiks cubes: %w", err)
	}

	return nil
}
