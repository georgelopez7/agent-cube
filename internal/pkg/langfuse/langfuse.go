package langfuse

import (
	"context"
	"log"

	"github.com/achetronic/adk-utils-go/plugin/langfuse"
	"google.golang.org/adk/runner"
)

type Langfuse struct {
	Config   runner.PluginConfig
	Shutdown func(context.Context) error
}

func NewLangfuse(ctx context.Context, name string, host string, publicKey string, secretKey string) *Langfuse {
	config, shutdown, err := langfuse.Setup(&langfuse.Config{
		ServiceName: name,
		Host:        host,
		PublicKey:   publicKey,
		SecretKey:   secretKey,
	})

	if err != nil {
		log.Fatalf("Failed to setup langfuse: %v", err)
	}

	return &Langfuse{
		Config:   config,
		Shutdown: shutdown,
	}
}
