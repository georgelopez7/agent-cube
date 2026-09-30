package mongo

import (
	"context"
	"log"

	"github.com/testcontainers/testcontainers-go/modules/mongodb"
)

func SetupMockMongo() (string, func()) {
	ctx := context.Background()

	container, err := mongodb.Run(ctx, "mongo:8")
	if err != nil {
		log.Fatalf("failed to start mock mongodb container: %v", err)
	}

	teardown := func() {
		if err := container.Terminate(ctx); err != nil {
			log.Fatalf("failed to terminate mock mongodb container: %v", err)
		}
	}

	uri, err := container.ConnectionString(ctx)
	if err != nil {
		log.Fatalf("failed to get mock mongodb connection string: %v", err)
	}

	return uri, teardown
}
